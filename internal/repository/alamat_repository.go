package repository

import (
	"github.com/Malikalfauziy20/evermos-miniproject/internal/entity"
	"gorm.io/gorm"
)

type AlamatRepository struct {
	DB *gorm.DB
}

func NewAlamatRepository(db *gorm.DB) *AlamatRepository {
	return &AlamatRepository{DB: db}
}

func (r *AlamatRepository) Create(alamat *entity.Alamat) error {
	return r.DB.Create(alamat).Error
}

func (r *AlamatRepository) FindAllByUserID(userID uint) ([]entity.Alamat, error) {
	var alamat []entity.Alamat
	err := r.DB.Where("id_user = ?", userID).Find(&alamat).Error
	return alamat, err
}

func (r *AlamatRepository) FindByID(id uint) (*entity.Alamat, error) {
	var alamat entity.Alamat
	err := r.DB.First(&alamat, id).Error
	if err != nil {
		return nil, err
	}
	return &alamat, nil
}

func (r *AlamatRepository) Update(alamat *entity.Alamat) error {
	return r.DB.Save(alamat).Error
}

func (r *AlamatRepository) Delete(id uint) error {
	return r.DB.Delete(&entity.Alamat{}, id).Error
}
