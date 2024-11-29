package handler

import (
	"ecommerce-app/internal/model"
	"ecommerce-app/internal/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate_userRoleMap = validator.New()

func GetUsersRoleMap(c *fiber.Ctx) error {
	userRolesMap, err := service.FetchUsersRoleMap()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return c.JSON(userRolesMap)
}

func CreateUserRoleMap(c *fiber.Ctx) error {
	var userRoleMap model.UserRoleMap
	if err := c.BodyParser(&userRoleMap); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	if err := validate_userRoleMap.Struct(&userRoleMap); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	if err := service.InsertUserRoleMap(&userRoleMap); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success"})
}

func UpdateUserRoleMap(c *fiber.Ctx) error {
	id := c.Params("id")
	var userRoleMap model.UserRoleMap
	if err := c.BodyParser(&userRoleMap); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	if err := validate_userRoleMap.Struct(&userRoleMap); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	if err := service.UpdateUserRoleMap(id, &userRoleMap); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success"})
}

func DeleteUserRoleMap(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := service.DeleteUserRoleMap(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success"})
}
