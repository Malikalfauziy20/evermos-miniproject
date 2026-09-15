package usecase

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Malikalfauziy20/evermos-miniproject/internal/dto"
	"github.com/Malikalfauziy20/evermos-miniproject/internal/entity"
	"github.com/Malikalfauziy20/evermos-miniproject/internal/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransaksiUsecase struct {
	DB            *gorm.DB
	TransaksiRepo *repository.TransaksiRepository
	AlamatRepo    *repository.AlamatRepository
	ProdukRepo    *repository.ProdukRepository
}

func NewTransaksiUsecase(db *gorm.DB, transaksiRepo *repository.TransaksiRepository, alamatRepo *repository.AlamatRepository, produkRepo *repository.ProdukRepository) *TransaksiUsecase {
	return &TransaksiUsecase{DB: db, TransaksiRepo: transaksiRepo, AlamatRepo: alamatRepo, ProdukRepo: produkRepo}
}

func (u *TransaksiUsecase) Checkout(userID uint, req dto.CheckoutRequest) (*entity.Transaksi, error) {
	alamat, err := u.AlamatRepo.FindByID(req.IDAlamat)
	if err != nil {
		return nil, errors.New("alamat tidak ditemukan")
	}
	if alamat.IDUser != userID {
		return nil, errors.New("alamat ini bukan milik kamu")
	}

	var trx entity.Transaksi

	err = u.DB.Transaction(func(tx *gorm.DB) error {
		trx = entity.Transaksi{
			IDUser:      userID,
			AlamatKirim: req.IDAlamat,
			KodeInvoice: "INV-" + uuid.NewString(),
			MethodBayar: req.MethodBayar,
		}
		if err := tx.Create(&trx).Error; err != nil {
			return err
		}

		totalHarga := 0

		for _, item := range req.DetailTrx {
			var produk entity.Produk
			if err := tx.First(&produk, item.IDProduk).Error; err != nil {
				return fmt.Errorf("produk id %d tidak ditemukan", item.IDProduk)
			}

			if produk.Stok < item.Kuantitas {
				return fmt.Errorf("stok produk %s tidak cukup", produk.NamaProduk)
			}

			log := entity.LogProduk{
				IDProduk:      produk.ID,
				NamaProduk:    produk.NamaProduk,
				Slug:          produk.Slug,
				HargaReseller: produk.HargaReseller,
				HargaKonsumen: produk.HargaKonsumen,
				Deskripsi:     produk.Deskripsi,
				IDToko:        produk.IDToko,
				IDCategory:    produk.IDCategory,
			}
			if err := tx.Create(&log).Error; err != nil {
				return err
			}

			detail := entity.DetailTransaksi{
				IDTransaksi: trx.ID,
				IDLogProduk: log.ID,
				IDToko:      produk.IDToko,
				Kuantitas:   item.Kuantitas,
			}
			if err := tx.Create(&detail).Error; err != nil {
				return err
			}

			if err := tx.Model(&produk).Update("stok", produk.Stok-item.Kuantitas).Error; err != nil {
				return err
			}

			hargaSatuan, convErr := strconv.Atoi(produk.HargaKonsumen)
			if convErr != nil {
				hargaSatuan = 0
			}
			totalHarga += hargaSatuan * item.Kuantitas
		}

		return tx.Model(&trx).Update("harga_total", totalHarga).Error
	})

	if err != nil {
		return nil, err
	}

	return u.TransaksiRepo.FindByID(trx.ID)
}

func (u *TransaksiUsecase) GetAll(userID uint) ([]entity.Transaksi, error) {
	return u.TransaksiRepo.FindAllByUserID(userID)
}

func (u *TransaksiUsecase) GetByID(userID, trxID uint) (*entity.Transaksi, error) {
	trx, err := u.TransaksiRepo.FindByID(trxID)
	if err != nil {
		return nil, err
	}
	if trx.IDUser != userID {
		return nil, errors.New("kamu tidak punya akses ke transaksi ini")
	}
	return trx, nil
}
