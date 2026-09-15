package usecase

import (
	"errors"

	"github.com/Malikalfauziy20/evermos-miniproject/internal/dto"
	"github.com/Malikalfauziy20/evermos-miniproject/internal/entity"
	"github.com/Malikalfauziy20/evermos-miniproject/internal/repository"
	"github.com/Malikalfauziy20/evermos-miniproject/pkg/bcrypt"
	appjwt "github.com/Malikalfauziy20/evermos-miniproject/pkg/jwt"
	"gorm.io/gorm"
)

type AuthUsecase struct {
	DB       *gorm.DB
	UserRepo *repository.UserRepository
}

func NewAuthUsecase(db *gorm.DB, userRepo *repository.UserRepository) *AuthUsecase {
	return &AuthUsecase{DB: db, UserRepo: userRepo}
}

func (u *AuthUsecase) Register(req dto.RegisterRequest) (*entity.User, error) {
	existingEmail, err := u.UserRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	if existingEmail != nil {
		return nil, errors.New("email sudah terdaftar")
	}

	existingPhone, err := u.UserRepo.FindByNoTelepon(req.NoTelepon)
	if err != nil {
		return nil, err
	}
	if existingPhone != nil {
		return nil, errors.New("no telepon sudah terdaftar")
	}

	hashedPassword, err := bcrypt.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := entity.User{
		Nama:       req.Nama,
		Email:      req.Email,
		NoTelepon:  req.NoTelepon,
		Password:   hashedPassword,
		Pekerjaan:  req.Pekerjaan,
		IDProvinsi: req.IDProvinsi,
		IDKota:     req.IDKota,
	}

	err = u.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			return err
		}

		toko := entity.Toko{
			IDUser:   user.ID,
			NamaToko: "Toko " + user.Nama,
		}
		return tx.Create(&toko).Error
	})

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *AuthUsecase) Login(req dto.LoginRequest) (string, *entity.User, error) {
	user, err := u.UserRepo.FindByNoTelepon(req.NoTelepon)
	if err != nil {
		return "", nil, err
	}
	if user == nil {
		return "", nil, errors.New("no telepon atau kata sandi salah")
	}

	if !bcrypt.CheckPassword(user.Password, req.Password) {
		return "", nil, errors.New("no telepon atau kata sandi salah")
	}

	token, err := appjwt.GenerateToken(user.ID, user.IsAdmin)
	if err != nil {
		return "", nil, err
	}
	return token, user, nil
}
