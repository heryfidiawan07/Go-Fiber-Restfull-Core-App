package controler

import "github.com/gofiber/fiber/v2"

// Hello hanlde api status
func Hello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": true, "message": "Hello i'm ok!", "data": nil})
}
