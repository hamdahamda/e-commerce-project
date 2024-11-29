package handler

import (
	"ecommerce-app/internal/model"
	"ecommerce-app/internal/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func GetActivities(c *fiber.Ctx) error {
	activities, err := service.FetchActivities()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return c.JSON(activities)
}

func CreateActivity(c *fiber.Ctx) error {
	var activity model.Activity
	if err := c.BodyParser(&activity); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	if err := validate.Struct(&activity); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	if err := service.InsertActivity(&activity); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success"})
}

func UpdateActivity(c *fiber.Ctx) error {
	id := c.Params("id")
	var activity model.Activity
	if err := c.BodyParser(&activity); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	if err := validate.Struct(&activity); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	if err := service.UpdateActivity(id, &activity); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success"})
}

func DeleteActivity(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := service.DeleteActivity(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success"})
}
