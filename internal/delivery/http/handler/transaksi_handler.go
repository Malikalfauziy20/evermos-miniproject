package handler

import (
	"strconv"

	"github.com/Malikalfauziy20/evermos-miniproject/internal/dto"
	"github.com/Malikalfauziy20/evermos-miniproject/internal/usecase"
	"github.com/Malikalfauziy20/evermos-miniproject/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type TransaksiHandler struct {
	TransaksiUsecase *usecase.TransaksiUsecase
}

func NewTransaksiHandler(transaksiUsecase *usecase.TransaksiUsecase) *TransaksiHandler {
	return &TransaksiHandler{TransaksiUsecase: transaksiUsecase}
}

func (h *TransaksiHandler) Checkout(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var req dto.CheckoutRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "body request tidak valid")
	}

	trx, err := h.TransaksiUsecase.Checkout(userID, req)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}
	return response.Created(c, "transaksi berhasil dibuat", trx)
}

func (h *TransaksiHandler) GetAll(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	trx, err := h.TransaksiUsecase.GetAll(userID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}
	return response.Success(c, "berhasil ambil semua transaksi", trx)
}

func (h *TransaksiHandler) GetByID(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "id tidak valid")
	}

	trx, err := h.TransaksiUsecase.GetByID(userID, uint(id))
	if err != nil {
		return response.Error(c, fiber.StatusForbidden, err.Error())
	}
	return response.Success(c, "berhasil ambil transaksi", trx)
}
