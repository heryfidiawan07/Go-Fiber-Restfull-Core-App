package controler

import (
	"api-fiber-gorm/request"
	"api-fiber-gorm/service"

	"github.com/gofiber/fiber/v2"
)

func RoleIndex(c *fiber.Ctx) error {
	return service.RoleIndex(c)
}

func RoleStore(c *fiber.Ctx) error {
	payload := new(request.RoleStore)

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": err, "data": nil})
	}

	if len(payload.Permissions) == 0 {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": "Permissions is required !", "data": nil})
	}

	err := request.RoleStoreValidation(*payload)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": err, "data": nil})
	}

	return service.RoleCreate(c, payload)
}

func RoleUpdate(c *fiber.Ctx) error {
	payload := new(request.RoleUpdate)

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": err, "data": nil})
	}

	if len(payload.Permissions) == 0 {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": "Permissions is required !", "data": nil})
	}

	err := request.RoleUpdateValidation(*payload)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": err, "data": nil})
	}

	return service.RoleEdit(c, payload)
}

func RoleShow(c *fiber.Ctx) error {
	return service.RoleFind(c)
}

func RoleDestroy(c *fiber.Ctx) error {
	return service.RoleDelete(c)
}
