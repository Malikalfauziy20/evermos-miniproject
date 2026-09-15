package usecase

import (
	"github.com/Malikalfauziy20/evermos-miniproject/internal/dto"
	"github.com/Malikalfauziy20/evermos-miniproject/internal/entity"
	"github.com/Malikalfauziy20/evermos-miniproject/internal/repository"
)

type KategoriUsecase struct {
	KategoriRepo *repository.KategoriRepository
}

func NewKategoriUsecase(kategoriRepo *repository.KategoriRepository) *KategoriUsecase {
	return &KategoriUsecase{KategoriRepo: kategoriRepo}
}

func (u *KategoriUsecase) Create(req dto.CreateKategoriRequest) (*entity.Kategori, error) {
	kategori := entity.Kategori{NamaKategori: req.NamaKategori}
	if err := u.KategoriRepo.Create(&kategori); err != nil {
		return nil, err
	}
	return &kategori, nil
}

func (u *KategoriUsecase) GetAll() ([]entity.Kategori, error) {
	return u.KategoriRepo.FindAll()
}

func (u *KategoriUsecase) GetByID(id uint) (*entity.Kategori, error) {
	return u.KategoriRepo.FindByID(id)
}

func (u *KategoriUsecase) Update(id uint, req dto.UpdateKategoriRequest) (*entity.Kategori, error) {
	kategori, err := u.KategoriRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	kategori.NamaKategori = req.NamaKategori
	if err := u.KategoriRepo.Update(kategori); err != nil {
		return nil, err
	}
	return kategori, nil
}

func (u *KategoriUsecase) Delete(id uint) error {
	return u.KategoriRepo.Delete(id)
}
