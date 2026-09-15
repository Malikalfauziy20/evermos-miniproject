package handler

import (
	"strconv"

	"github.com/Malikalfauziy20/evermos-miniproject/internal/dto"
	"github.com/Malikalfauziy20/evermos-miniproject/internal/usecase"
	"github.com/Malikalfauziy20/evermos-miniproject/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type ProdukHandler struct {
	ProdukUsecase *usecase.ProdukUsecase
}

func NewProdukHandler(produkUsecase *usecase.ProdukUsecase) *ProdukHandler {
	return &ProdukHandler{ProdukUsecase: produkUsecase}
}

func (h *ProdukHandler) Create(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var req dto.CreateProdukRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "body request tidak valid")
	}

	produk, err := h.ProdukUsecase.Create(userID, req)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}
	return response.Created(c, "produk berhasil dibuat", produk)
}

func (h *ProdukHandler) GetAll(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	categoryID, _ := strconv.Atoi(c.Query("category_id", "0"))

	filter := dto.ProdukFilter{
		Page:       page,
		Limit:      limit,
		NamaProduk: c.Query("nama_produk", ""),
		CategoryID: uint(categoryID),
	}

	produk, total, err := h.ProdukUsecase.GetAll(filter)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "berhasil ambil semua produk",
		"data":    produk,
		"meta": fiber.Map{
			"page":       filter.Page,
			"limit":      filter.Limit,
			"total_data": total,
			"total_page": (total + int64(filter.Limit) - 1) / int64(filter.Limit),
		},
	})
}

func (h *ProdukHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "id tidak valid")
	}

	produk, err := h.ProdukUsecase.GetByID(uint(id))
	if err != nil {
		return response.Error(c, fiber.StatusNotFound, "produk tidak ditemukan")
	}
	return response.Success(c, "berhasil ambil produk", produk)
}

func (h *ProdukHandler) Update(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "id tidak valid")
	}

	var req dto.UpdateProdukRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "body request tidak valid")
	}

	produk, err := h.ProdukUsecase.Update(userID, uint(id), req)
	if err != nil {
		return response.Error(c, fiber.StatusForbidden, err.Error())
	}
	return response.Success(c, "produk berhasil diupdate", produk)
}

func (h *ProdukHandler) Delete(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "id tidak valid")
	}

	if err := h.ProdukUsecase.Delete(userID, uint(id)); err != nil {
		return response.Error(c, fiber.StatusForbidden, err.Error())
	}
	return response.Success(c, "produk berhasil dihapus", nil)
}

func (h *ProdukHandler) UploadFoto(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "id tidak valid")
	}

	file, err := c.FormFile("foto")
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "file foto wajib diisi")
	}

	foto, err := h.ProdukUsecase.UploadFoto(c, userID, uint(id), file)
	if err != nil {
		return response.Error(c, fiber.StatusForbidden, err.Error())
	}
	return response.Created(c, "foto produk berhasil diupload", foto)
}
