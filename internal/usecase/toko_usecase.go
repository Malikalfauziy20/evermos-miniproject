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

type TokoUsecase struct {
	TokoRepo *repository.TokoRepository
}

func NewTokoUsecase(tokoRepo *repository.TokoRepository) *TokoUsecase {
	return &TokoUsecase{TokoRepo: tokoRepo}
}

func (u *TokoUsecase) GetMyToko(userID uint) (*entity.Toko, error) {
	return u.TokoRepo.FindByUserID(userID)
}

func (u *TokoUsecase) GetByID(id uint) (*entity.Toko, error) {
	return u.TokoRepo.FindByID(id)
}

func (u *TokoUsecase) Update(c *fiber.Ctx, userID uint, tokoID uint, req dto.UpdateTokoRequest, file *multipart.FileHeader) (*entity.Toko, error) {
	toko, err := u.TokoRepo.FindByID(tokoID)
	if err != nil {
		return nil, err
	}
	if toko.IDUser != userID {
		return nil, errors.New("kamu tidak punya akses ke toko ini")
	}

	if req.NamaToko != "" {
		toko.NamaToko = req.NamaToko
	}

	if file != nil {
		path, err := upload.SaveFile(c, file, "toko")
		if err != nil {
			return nil, err
		}
		toko.URLFoto = path
	}

	if err := u.TokoRepo.Update(toko); err != nil {
		return nil, err
	}
	return toko, nil
}
