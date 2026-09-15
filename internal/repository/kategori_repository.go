package repository

import (
	"github.com/Malikalfauziy20/evermos-miniproject/internal/entity"
	"gorm.io/gorm"
)

type KategoriRepository struct {
	DB *gorm.DB
}

func NewKategoriRepository(db *gorm.DB) *KategoriRepository {
	return &KategoriRepository{DB: db}
}

func (r *KategoriRepository) Create(kategori *entity.Kategori) error {
	return r.DB.Create(kategori).Error
}

func (r *KategoriRepository) FindAll() ([]entity.Kategori, error) {
	var kategori []entity.Kategori
	err := r.DB.Find(&kategori).Error
	return kategori, err
}

func (r *KategoriRepository) FindByID(id uint) (*entity.Kategori, error) {
	var kategori entity.Kategori
	err := r.DB.First(&kategori, id).Error
	if err != nil {
		return nil, err
	}
	return &kategori, nil
}

func (r *KategoriRepository) Update(kategori *entity.Kategori) error {
	return r.DB.Save(kategori).Error
}

func (r *KategoriRepository) Delete(id uint) error {
	return r.DB.Delete(&entity.Kategori{}, id).Error
}
