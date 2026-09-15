package handler

import (
	"strconv"

	"github.com/Malikalfauziy20/evermos-miniproject/internal/dto"
	"github.com/Malikalfauziy20/evermos-miniproject/internal/usecase"
	"github.com/Malikalfauziy20/evermos-miniproject/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type AlamatHandler struct {
	AlamatUsecase *usecase.AlamatUsecase
}

func NewAlamatHandler(alamatUsecase *usecase.AlamatUsecase) *AlamatHandler {
	return &AlamatHandler{AlamatUsecase: alamatUsecase}
}

func (h *AlamatHandler) Create(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var req dto.CreateAlamatRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "body request tidak valid")
	}

	alamat, err := h.AlamatUsecase.Create(userID, req)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}
	return response.Created(c, "alamat berhasil dibuat", alamat)
}

func (h *AlamatHandler) GetAll(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	alamat, err := h.AlamatUsecase.GetAll(userID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}
	return response.Success(c, "berhasil ambil semua alamat", alamat)
}

func (h *AlamatHandler) GetByID(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "id tidak valid")
	}

	alamat, err := h.AlamatUsecase.GetByID(userID, uint(id))
	if err != nil {
		return response.Error(c, fiber.StatusForbidden, err.Error())
	}
	return response.Success(c, "berhasil ambil alamat", alamat)
}

func (h *AlamatHandler) Update(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "id tidak valid")
	}

	var req dto.UpdateAlamatRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "body request tidak valid")
	}

	alamat, err := h.AlamatUsecase.Update(userID, uint(id), req)
	if err != nil {
		return response.Error(c, fiber.StatusForbidden, err.Error())
	}
	return response.Success(c, "alamat berhasil diupdate", alamat)
}

func (h *AlamatHandler) Delete(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "id tidak valid")
	}

	if err := h.AlamatUsecase.Delete(userID, uint(id)); err != nil {
		return response.Error(c, fiber.StatusForbidden, err.Error())
	}
	return response.Success(c, "alamat berhasil dihapus", nil)
}
