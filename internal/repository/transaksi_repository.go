package repository

import (
	"github.com/Malikalfauziy20/evermos-miniproject/internal/entity"
	"gorm.io/gorm"
)

type TransaksiRepository struct {
	DB *gorm.DB
}

func NewTransaksiRepository(db *gorm.DB) *TransaksiRepository {
	return &TransaksiRepository{DB: db}
}

func (r *TransaksiRepository) FindAllByUserID(userID uint) ([]entity.Transaksi, error) {
	var trx []entity.Transaksi
	err := r.DB.Preload("DetailTransaksi").
		Where("id_user = ?", userID).Find(&trx).Error
	return trx, err
}

func (r *TransaksiRepository) FindByID(id uint) (*entity.Transaksi, error) {
	var trx entity.Transaksi
	err := r.DB.Preload("DetailTransaksi.LogProduk").First(&trx, id).Error
	if err != nil {
		return nil, err
	}
	return &trx, nil
}
