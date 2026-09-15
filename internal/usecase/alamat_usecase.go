package usecase

import (
	"errors"

	"github.com/Malikalfauziy20/evermos-miniproject/internal/dto"
	"github.com/Malikalfauziy20/evermos-miniproject/internal/entity"
	"github.com/Malikalfauziy20/evermos-miniproject/internal/repository"
)

type AlamatUsecase struct {
	AlamatRepo *repository.AlamatRepository
}

func NewAlamatUsecase(alamatRepo *repository.AlamatRepository) *AlamatUsecase {
	return &AlamatUsecase{AlamatRepo: alamatRepo}
}

func (u *AlamatUsecase) Create(userID uint, req dto.CreateAlamatRequest) (*entity.Alamat, error) {
	alamat := entity.Alamat{
		IDUser:       userID,
		JudulAlamat:  req.JudulAlamat,
		NamaPenerima: req.NamaPenerima,
		NoTelp:       req.NoTelp,
		DetailAlamat: req.DetailAlamat,
	}
	if err := u.AlamatRepo.Create(&alamat); err != nil {
		return nil, err
	}
	return &alamat, nil
}

func (u *AlamatUsecase) GetAll(userID uint) ([]entity.Alamat, error) {
	return u.AlamatRepo.FindAllByUserID(userID)
}

func (u *AlamatUsecase) checkOwnership(userID, alamatID uint) (*entity.Alamat, error) {
	alamat, err := u.AlamatRepo.FindByID(alamatID)
	if err != nil {
		return nil, err
	}
	if alamat.IDUser != userID {
		return nil, errors.New("kamu tidak punya akses ke alamat ini")
	}
	return alamat, nil
}

func (u *AlamatUsecase) GetByID(userID, alamatID uint) (*entity.Alamat, error) {
	return u.checkOwnership(userID, alamatID)
}

func (u *AlamatUsecase) Update(userID, alamatID uint, req dto.UpdateAlamatRequest) (*entity.Alamat, error) {
	alamat, err := u.checkOwnership(userID, alamatID)
	if err != nil {
		return nil, err
	}

	if req.JudulAlamat != "" {
		alamat.JudulAlamat = req.JudulAlamat
	}
	if req.NamaPenerima != "" {
		alamat.NamaPenerima = req.NamaPenerima
	}
	if req.NoTelp != "" {
		alamat.NoTelp = req.NoTelp
	}
	if req.DetailAlamat != "" {
		alamat.DetailAlamat = req.DetailAlamat
	}

	if err := u.AlamatRepo.Update(alamat); err != nil {
		return nil, err
	}
	return alamat, nil
}

func (u *AlamatUsecase) Delete(userID, alamatID uint) error {
	_, err := u.checkOwnership(userID, alamatID)
	if err != nil {
		return err
	}
	return u.AlamatRepo.Delete(alamatID)
}
