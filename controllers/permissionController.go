package controllers

import (
	"github.com/gofiber/fiber"
	"goreact/database"
	"goreact/models"
)

func AllPermissions(c *fiber.Ctx) error {
	var permissions []models.Permission

	database.DB.Find(&permissions)

	return c.JSON(permissions)
}
