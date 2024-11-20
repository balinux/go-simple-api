package handler

import (
	"go-fiber-hairstyle/config"
	"go-fiber-hairstyle/db"
	"go-fiber-hairstyle/routes"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	r.RequestURI = r.URL.String()
	handler().ServeHTTP(w, r)
}

func handler() http.HandlerFunc {
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

	return adaptor.FiberApp(app)
}
