package hairstyles

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type HairStyleHandler struct {
	service *HairStyleService
}

func NewHairstyleHandler(service *HairStyleService) *HairStyleHandler {
	return &HairStyleHandler{service: service}
}

func (h HairStyleHandler) GetAll(c *fiber.Ctx) error {
	hairStyle, err := h.service.GetAllHairStyles()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch hair styles"})
	}
	return c.JSON(hairStyle)
}

func (h HairStyleHandler) Create(c *fiber.Ctx) error {
	var newHairStyle HairStyle
	if err := c.BodyParser(&newHairStyle); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	if err := h.service.CreateHairStyle(&newHairStyle); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": " Failed to create hairstyles",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(newHairStyle)
}

func (h HairStyleHandler) GetHairStyleByID(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	hairStyele, err := h.service.GetHairStyleByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "hairstyles not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(hairStyele)
}

func (h *HairStyleHandler) DeleteHairStyleByID(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID parameter",
		})
	}

	if err := h.service.DeleteHairStyleByID(uint(id)); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Hairstyle not found",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *HairStyleHandler) UpdateHairStyle(c *fiber.Ctx) error {
	// ambil id berdasarkan parameter id
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid ID parameter",
		})
	}

	// cari data dengan id
	var hairStyle HairStyle
	if err := c.BodyParser(&hairStyle); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	// update data dengan id yang sudah di dapatkan
	// memasukkan id dengan id yang sesuai
	hairStyle.ID = uint(id)

	if err := h.service.UpdateHairStyle(&hairStyle); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed go update hairstyles",
		})
	}
	return c.JSON(hairStyle)
}
