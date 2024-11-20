package routes

import (
	"go-fiber-hairstyle/db"
	"go-fiber-hairstyle/internal/hairstyles"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	hairStyleRepo := hairstyles.NewHairStyleRepository(db.DB)
	hairStyleService := hairstyles.NewHairStyleService(hairStyleRepo)
	hairStyleHandler := hairstyles.NewHairstyleHandler(hairStyleService)

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/hairstyles", hairStyleHandler.GetAll)
	v1.Post("/hairstyles", hairStyleHandler.Create)
	v1.Get("/hairstyles/:id", hairStyleHandler.GetHairStyleByID)
	v1.Delete("/hairstyles/:id", hairStyleHandler.DeleteHairStyleByID)
	v1.Put("/hairstyles/:id", hairStyleHandler.UpdateHairStyle)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "hairstyles api is ok",
		})
	})

	// uncomment if you need metrics
	// app.Get("/metrics", monitor.New())
}
