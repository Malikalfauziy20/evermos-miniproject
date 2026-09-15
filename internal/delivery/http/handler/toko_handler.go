package handler

import (
	"strconv"

	"github.com/Malikalfauziy20/evermos-miniproject/internal/dto"
	"github.com/Malikalfauziy20/evermos-miniproject/internal/usecase"
	"github.com/Malikalfauziy20/evermos-miniproject/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type TokoHandler struct {
	TokoUsecase *usecase.TokoUsecase
}

func NewTokoHandler(tokoUsecase *usecase.TokoUsecase) *TokoHandler {
	return &TokoHandler{TokoUsecase: tokoUsecase}
}

func (h *TokoHandler) GetMyToko(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	toko, err := h.TokoUsecase.GetMyToko(userID)
	if err != nil {
		return response.Error(c, fiber.StatusNotFound, "toko tidak ditemukan")
	}
	return response.Success(c, "berhasil ambil data toko", toko)
}

func (h *TokoHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "id tidak valid")
	}

	toko, err := h.TokoUsecase.GetByID(uint(id))
	if err != nil {
		return response.Error(c, fiber.StatusNotFound, "toko tidak ditemukan")
	}
	return response.Success(c, "berhasil ambil data toko", toko)
}

func (h *TokoHandler) Update(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "id tidak valid")
	}

	var req dto.UpdateTokoRequest
	req.NamaToko = c.FormValue("nama_toko")

	file, _ := c.FormFile("foto")

	toko, err := h.TokoUsecase.Update(c, userID, uint(id), req, file)
	if err != nil {
		return response.Error(c, fiber.StatusForbidden, err.Error())
	}
	return response.Success(c, "toko berhasil diupdate", toko)
}
