package main

import (
	"go-fiber-hairstyle/config"
	"go-fiber-hairstyle/db"
	"go-fiber-hairstyle/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Inisialisasi konfigurasi
	config.LoadConfig()

	// Inisialisasi database
	db.Init()

	// Inisialisasi Fiber
	app := fiber.New()

	// middleware logger
	app.Use(logger.New())

	// Daftar routes
	routes.Setup(app)

	// Jalankan aplikasi
	log.Fatal(app.Listen(":3000"))
}
