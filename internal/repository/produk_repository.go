package repository

import (
	"github.com/Malikalfauziy20/evermos-miniproject/internal/dto"
	"github.com/Malikalfauziy20/evermos-miniproject/internal/entity"
	"gorm.io/gorm"
)

type ProdukRepository struct {
	DB *gorm.DB
}

func NewProdukRepository(db *gorm.DB) *ProdukRepository {
	return &ProdukRepository{DB: db}
}

func (r *ProdukRepository) Create(produk *entity.Produk) error {
	return r.DB.Create(produk).Error
}

func (r *ProdukRepository) FindAll(filter dto.ProdukFilter) ([]entity.Produk, int64, error) {
	var produk []entity.Produk
	var total int64

	query := r.DB.Model(&entity.Produk{})

	if filter.NamaProduk != "" {
		query = query.Where("nama_produk LIKE ?", "%"+filter.NamaProduk+"%")
	}
	if filter.CategoryID != 0 {
		query = query.Where("id_category = ?", filter.CategoryID)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (filter.Page - 1) * filter.Limit
	err := query.Preload("Foto").Preload("Kategori").Preload("Toko").
		Offset(offset).Limit(filter.Limit).Find(&produk).Error

	return produk, total, err
}

func (r *ProdukRepository) FindByID(id uint) (*entity.Produk, error) {
	var produk entity.Produk
	err := r.DB.Preload("Foto").Preload("Kategori").Preload("Toko").First(&produk, id).Error
	if err != nil {
		return nil, err
	}
	return &produk, nil
}

func (r *ProdukRepository) Update(produk *entity.Produk) error {
	return r.DB.Save(produk).Error
}

func (r *ProdukRepository) Delete(id uint) error {
	return r.DB.Delete(&entity.Produk{}, id).Error
}

func (r *ProdukRepository) CreateFoto(foto *entity.FotoProduk) error {
	return r.DB.Create(foto).Error
}
