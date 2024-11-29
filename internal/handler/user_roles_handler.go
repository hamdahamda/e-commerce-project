package handler

import (
	"ecommerce-app/internal/model"
	"ecommerce-app/internal/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate_userRole = validator.New()

func GetUsersRole(c *fiber.Ctx) error {
	userRoles, err := service.FetchUsersRole()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return c.JSON(userRoles)
}

func CreateUserRole(c *fiber.Ctx) error {
	var userRole model.UserRole
	if err := c.BodyParser(&userRole); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	if err := validate_userRole.Struct(&userRole); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	if err := service.InsertUserRole(&userRole); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success"})
}

func UpdateUserRole(c *fiber.Ctx) error {
	id := c.Params("id")
	var userRole model.UserRole
	if err := c.BodyParser(&userRole); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	if err := validate_userRole.Struct(&userRole); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	if err := service.UpdateUserRole(id, &userRole); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success"})
}

func DeleteUserRole(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := service.DeleteUserRole(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success"})
}
