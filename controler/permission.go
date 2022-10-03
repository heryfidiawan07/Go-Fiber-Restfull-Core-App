package controler

import (
	"api-fiber-gorm/service"

	"github.com/gofiber/fiber/v2"
)

func PermissionIndex(c *fiber.Ctx) error {
	return service.PermissionIndex(c)
}
