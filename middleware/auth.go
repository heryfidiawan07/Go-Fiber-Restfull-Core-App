package middleware

import (
	"api-fiber-gorm/config"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func Auth() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(config.Config("SECRET")),
		ErrorHandler: AuthJwtError,
	})
}

func AuthJwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(400).
			JSON(fiber.Map{"status": false, "message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(401).
		JSON(fiber.Map{"status": false, "message": "Invalid or expired JWT", "data": nil})
}
