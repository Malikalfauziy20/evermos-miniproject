package usecase

import (
	"github.com/Malikalfauziy20/evermos-miniproject/internal/dto"
	"github.com/Malikalfauziy20/evermos-miniproject/internal/entity"
	"github.com/Malikalfauziy20/evermos-miniproject/internal/repository"
)

type UserUsecase struct {
	UserRepo *repository.UserRepository
}

func NewUserUsecase(userRepo *repository.UserRepository) *UserUsecase {
	return &UserUsecase{UserRepo: userRepo}
}

func (u *UserUsecase) GetProfile(userID uint) (*entity.User, error) {
	return u.UserRepo.FindByID(userID)
}

func (u *UserUsecase) UpdateProfile(userID uint, req dto.UpdateUserRequest) (*entity.User, error) {
	user, err := u.UserRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	if req.Nama != "" {
		user.Nama = req.Nama
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.NoTelepon != "" {
		user.NoTelepon = req.NoTelepon
	}
	if req.Pekerjaan != "" {
		user.Pekerjaan = req.Pekerjaan
	}
	if req.IDProvinsi != "" {
		user.IDProvinsi = req.IDProvinsi
	}
	if req.IDKota != "" {
		user.IDKota = req.IDKota
	}

	if err := u.UserRepo.Update(user); err != nil {
		return nil, err
	}
	return user, nil
}
