package controler

import (
	"api-fiber-gorm/request"
	"api-fiber-gorm/service"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	payload := new(request.Login)

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": err, "data": nil})
	}

	err := request.LoginValidation(*payload)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": err, "data": nil})
	}

	return service.AuthLogin(c, payload)
}

func Register(c *fiber.Ctx) error {
	payload := new(request.Register)

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": err, "data": nil})
	}

	err := request.RegisterValidation(*payload)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": err, "data": nil})
	}

	return service.AuthRegister(c, payload)
}

func RefreshToken(c *fiber.Ctx) error {
	payload := new(request.RefreshToken)

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": err, "data": nil})
	}

	return service.AuthRefreshToken(c, payload)
}

func Me(c *fiber.Ctx) error {
	return service.AuthMe(c)
}
