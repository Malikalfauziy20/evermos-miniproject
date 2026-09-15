package main

import (
	"log"
	"os"

	"github.com/Malikalfauziy20/evermos-miniproject/config"
	"github.com/Malikalfauziy20/evermos-miniproject/internal/delivery/http/handler"
	"github.com/Malikalfauziy20/evermos-miniproject/internal/delivery/http/router"
	"github.com/Malikalfauziy20/evermos-miniproject/internal/entity"
	"github.com/Malikalfauziy20/evermos-miniproject/internal/repository"
	"github.com/Malikalfauziy20/evermos-miniproject/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("tidak menemukan file .env, lanjut pakai environment variable sistem")
	}

	db := config.ConnectDB()
	if err := db.AutoMigrate(
		&entity.User{}, &entity.Toko{}, &entity.Alamat{},
		&entity.Kategori{}, &entity.Produk{}, &entity.FotoProduk{},
		&entity.Transaksi{}, &entity.DetailTransaksi{}, &entity.LogProduk{},
	); err != nil {
		log.Fatal("migrasi gagal: ", err)
	}
	log.Println("migrasi berhasil, database siap")

	userRepo := repository.NewUserRepository(db)
	tokoRepo := repository.NewTokoRepository(db)
	alamatRepo := repository.NewAlamatRepository(db)
	kategoriRepo := repository.NewKategoriRepository(db)
	produkRepo := repository.NewProdukRepository(db)
	transaksiRepo := repository.NewTransaksiRepository(db)

	authUsecase := usecase.NewAuthUsecase(db, userRepo)
	userUsecase := usecase.NewUserUsecase(userRepo)
	tokoUsecase := usecase.NewTokoUsecase(tokoRepo)
	alamatUsecase := usecase.NewAlamatUsecase(alamatRepo)
	kategoriUsecase := usecase.NewKategoriUsecase(kategoriRepo)
	produkUsecase := usecase.NewProdukUsecase(produkRepo, tokoRepo)
	transaksiUsecase := usecase.NewTransaksiUsecase(db, transaksiRepo, alamatRepo, produkRepo)

	authHandler := handler.NewAuthHandler(authUsecase)
	userHandler := handler.NewUserHandler(userUsecase)
	tokoHandler := handler.NewTokoHandler(tokoUsecase)
	alamatHandler := handler.NewAlamatHandler(alamatUsecase)
	kategoriHandler := handler.NewKategoriHandler(kategoriUsecase)
	produkHandler := handler.NewProdukHandler(produkUsecase)
	transaksiHandler := handler.NewTransaksiHandler(transaksiUsecase)

	app := fiber.New()

	router.SetupRoutes(app, router.Handlers{
		Auth:      authHandler,
		User:      userHandler,
		Toko:      tokoHandler,
		Alamat:    alamatHandler,
		Kategori:  kategoriHandler,
		Produk:    produkHandler,
		Transaksi: transaksiHandler,
	})

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("server berjalan di port " + port)
	log.Fatal(app.Listen(":" + port))
}
