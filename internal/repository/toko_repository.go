package repository

import (
	"github.com/Malikalfauziy20/evermos-miniproject/internal/entity"
	"gorm.io/gorm"
)

type TokoRepository struct {
	DB *gorm.DB
}

func NewTokoRepository(db *gorm.DB) *TokoRepository {
	return &TokoRepository{DB: db}
}

func (r *TokoRepository) FindByUserID(userID uint) (*entity.Toko, error) {
	var toko entity.Toko
	err := r.DB.Where("id_user = ?", userID).First(&toko).Error
	if err != nil {
		return nil, err
	}
	return &toko, nil
}

func (r *TokoRepository) FindByID(id uint) (*entity.Toko, error) {
	var toko entity.Toko
	err := r.DB.First(&toko, id).Error
	if err != nil {
		return nil, err
	}
	return &toko, nil
}

func (r *TokoRepository) Update(toko *entity.Toko) error {
	return r.DB.Save(toko).Error
}
