package handler

import (
	"strconv"

	"github.com/Malikalfauziy20/evermos-miniproject/internal/dto"
	"github.com/Malikalfauziy20/evermos-miniproject/internal/usecase"
	"github.com/Malikalfauziy20/evermos-miniproject/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type KategoriHandler struct {
	KategoriUsecase *usecase.KategoriUsecase
}

func NewKategoriHandler(kategoriUsecase *usecase.KategoriUsecase) *KategoriHandler {
	return &KategoriHandler{KategoriUsecase: kategoriUsecase}
}

func (h *KategoriHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateKategoriRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "body request tidak valid")
	}

	kategori, err := h.KategoriUsecase.Create(req)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}
	return response.Created(c, "kategori berhasil dibuat", kategori)
}

func (h *KategoriHandler) GetAll(c *fiber.Ctx) error {
	kategori, err := h.KategoriUsecase.GetAll()
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}
	return response.Success(c, "berhasil ambil semua kategori", kategori)
}

func (h *KategoriHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "id tidak valid")
	}

	var req dto.UpdateKategoriRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "body request tidak valid")
	}

	kategori, err := h.KategoriUsecase.Update(uint(id), req)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}
	return response.Success(c, "kategori berhasil diupdate", kategori)
}

func (h *KategoriHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "id tidak valid")
	}

	if err := h.KategoriUsecase.Delete(uint(id)); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}
	return response.Success(c, "kategori berhasil dihapus", nil)
}
