package handler

import (
	"github.com/Malikalfauziy20/evermos-miniproject/internal/dto"
	"github.com/Malikalfauziy20/evermos-miniproject/internal/usecase"
	"github.com/Malikalfauziy20/evermos-miniproject/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserUsecase *usecase.UserUsecase
}

func NewUserHandler(userUsecase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{UserUsecase: userUsecase}
}

func (h *UserHandler) GetProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	user, err := h.UserUsecase.GetProfile(userID)
	if err != nil {
		return response.Error(c, fiber.StatusNotFound, "user tidak ditemukan")
	}
	return response.Success(c, "berhasil ambil profil", user)
}

func (h *UserHandler) UpdateProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var req dto.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "body request tidak valid")
	}

	user, err := h.UserUsecase.UpdateProfile(userID, req)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}
	return response.Success(c, "profil berhasil diupdate", user)
}
