package router

import (
	"github.com/Malikalfauziy20/evermos-miniproject/internal/delivery/http/handler"
	"github.com/Malikalfauziy20/evermos-miniproject/internal/delivery/http/middleware"
	"github.com/gofiber/fiber/v2"
)

type Handlers struct {
	Auth      *handler.AuthHandler
	User      *handler.UserHandler
	Toko      *handler.TokoHandler
	Alamat    *handler.AlamatHandler
	Kategori  *handler.KategoriHandler
	Produk    *handler.ProdukHandler
	Transaksi *handler.TransaksiHandler
}

func SetupRoutes(app *fiber.App, h Handlers) {
	auth := app.Group("/auth")
	auth.Post("/register", h.Auth.Register)
	auth.Post("/login", h.Auth.Login)

	user := app.Group("/user", middleware.JWTProtected())
	user.Get("/", h.User.GetProfile)
	user.Put("/", h.User.UpdateProfile)

	toko := app.Group("/toko", middleware.JWTProtected())
	toko.Get("/my", h.Toko.GetMyToko)
	toko.Get("/:id", h.Toko.GetByID)
	toko.Put("/:id", h.Toko.Update)

	alamat := app.Group("/alamat", middleware.JWTProtected())
	alamat.Post("/", h.Alamat.Create)
	alamat.Get("/", h.Alamat.GetAll)
	alamat.Get("/:id", h.Alamat.GetByID)
	alamat.Put("/:id", h.Alamat.Update)
	alamat.Delete("/:id", h.Alamat.Delete)

	kategori := app.Group("/kategori", middleware.JWTProtected())
	kategori.Get("/", h.Kategori.GetAll)
	kategori.Post("/", middleware.AdminOnly(), h.Kategori.Create)
	kategori.Put("/:id", middleware.AdminOnly(), h.Kategori.Update)
	kategori.Delete("/:id", middleware.AdminOnly(), h.Kategori.Delete)

	produk := app.Group("/produk", middleware.JWTProtected())
	produk.Post("/", h.Produk.Create)
	produk.Get("/", h.Produk.GetAll)
	produk.Get("/:id", h.Produk.GetByID)
	produk.Put("/:id", h.Produk.Update)
	produk.Delete("/:id", h.Produk.Delete)
	produk.Post("/:id/foto", h.Produk.UploadFoto)

	trx := app.Group("/trx", middleware.JWTProtected())
	trx.Post("/", h.Transaksi.Checkout)
	trx.Get("/", h.Transaksi.GetAll)
	trx.Get("/:id", h.Transaksi.GetByID)
}
