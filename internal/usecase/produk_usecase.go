package usecase

import (
	"errors"
	"mime/multipart"

	"github.com/Malikalfauziy20/evermos-miniproject/internal/dto"
	"github.com/Malikalfauziy20/evermos-miniproject/internal/entity"
	"github.com/Malikalfauziy20/evermos-miniproject/internal/repository"
	"github.com/Malikalfauziy20/evermos-miniproject/pkg/upload"
	"github.com/gofiber/fiber/v2"
)

type ProdukUsecase struct {
	ProdukRepo *repository.ProdukRepository
	TokoRepo   *repository.TokoRepository
}

func NewProdukUsecase(produkRepo *repository.ProdukRepository, tokoRepo *repository.TokoRepository) *ProdukUsecase {
	return &ProdukUsecase{ProdukRepo: produkRepo, TokoRepo: tokoRepo}
}

func (u *ProdukUsecase) Create(userID uint, req dto.CreateProdukRequest) (*entity.Produk, error) {
	toko, err := u.TokoRepo.FindByUserID(userID)
	if err != nil {
		return nil, errors.New("toko kamu tidak ditemukan")
	}

	produk := entity.Produk{
		NamaProduk:    req.NamaProduk,
		HargaReseller: req.HargaReseller,
		HargaKonsumen: req.HargaKonsumen,
		Stok:          req.Stok,
		Deskripsi:     req.Deskripsi,
		IDToko:        toko.ID,
		IDCategory:    req.IDCategory,
	}

	if err := u.ProdukRepo.Create(&produk); err != nil {
		return nil, err
	}
	return &produk, nil
}

func (u *ProdukUsecase) GetAll(filter dto.ProdukFilter) ([]entity.Produk, int64, error) {
	if filter.Page <= 0 {
		filter.Page = 1
	}
	if filter.Limit <= 0 {
		filter.Limit = 10
	}
	return u.ProdukRepo.FindAll(filter)
}

func (u *ProdukUsecase) GetByID(id uint) (*entity.Produk, error) {
	return u.ProdukRepo.FindByID(id)
}

func (u *ProdukUsecase) checkOwnership(userID, produkID uint) (*entity.Produk, error) {
	produk, err := u.ProdukRepo.FindByID(produkID)
	if err != nil {
		return nil, err
	}
	toko, err := u.TokoRepo.FindByID(produk.IDToko)
	if err != nil {
		return nil, err
	}
	if toko.IDUser != userID {
		return nil, errors.New("kamu tidak punya akses ke produk ini")
	}
	return produk, nil
}

func (u *ProdukUsecase) Update(userID, produkID uint, req dto.UpdateProdukRequest) (*entity.Produk, error) {
	produk, err := u.checkOwnership(userID, produkID)
	if err != nil {
		return nil, err
	}

	if req.NamaProduk != "" {
		produk.NamaProduk = req.NamaProduk
	}
	if req.HargaReseller != "" {
		produk.HargaReseller = req.HargaReseller
	}
	if req.HargaKonsumen != "" {
		produk.HargaKonsumen = req.HargaKonsumen
	}
	if req.Stok != nil {
		produk.Stok = *req.Stok
	}
	if req.Deskripsi != "" {
		produk.Deskripsi = req.Deskripsi
	}
	if req.IDCategory != 0 {
		produk.IDCategory = req.IDCategory
	}

	if err := u.ProdukRepo.Update(produk); err != nil {
		return nil, err
	}
	return produk, nil
}

func (u *ProdukUsecase) Delete(userID, produkID uint) error {
	_, err := u.checkOwnership(userID, produkID)
	if err != nil {
		return err
	}
	return u.ProdukRepo.Delete(produkID)
}

func (u *ProdukUsecase) UploadFoto(c *fiber.Ctx, userID, produkID uint, file *multipart.FileHeader) (*entity.FotoProduk, error) {
	_, err := u.checkOwnership(userID, produkID)
	if err != nil {
		return nil, err
	}

	path, err := upload.SaveFile(c, file, "produk")
	if err != nil {
		return nil, err
	}

	foto := entity.FotoProduk{IDProduk: produkID, URL: path}
	if err := u.ProdukRepo.CreateFoto(&foto); err != nil {
		return nil, err
	}
	return &foto, nil
}
