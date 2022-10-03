package service

import (
	"api-fiber-gorm/database"
	"api-fiber-gorm/response"

	"github.com/gofiber/fiber/v2"
)

func PermissionIndex(c *fiber.Ctx) error {
	var permissions []response.Permission

	db := database.DB
	if err := db.Find(&permissions).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": false, "message": err, "data": nil})
	}

	return c.JSON(fiber.Map{"status": true, "message": nil, "data": permissions})
}
