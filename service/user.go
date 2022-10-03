package service

import (
	"api-fiber-gorm/database"
	"api-fiber-gorm/helper"
	"api-fiber-gorm/model"
	"api-fiber-gorm/request"
	"api-fiber-gorm/response"

	"github.com/gofiber/fiber/v2"
)

func UserIndex(c *fiber.Ctx) error {
	var users []response.User

	db := database.DB
	if err := db.Find(&users).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": false, "message": err, "data": nil})
	}

	return c.JSON(fiber.Map{"status": true, "message": nil, "data": users})
}

func UserCreate(c *fiber.Ctx, request *request.UserStore) error {
	if !helper.ValidEmail(request.Email) {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": "Email not valid !", "data": nil})
	}

	password, err := helper.HashPassword(request.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": false, "message": err, "data": nil})
	}

	user := model.User{
		Name:     request.Name,
		Username: request.Username,
		Email:    request.Email,
		Password: password,
		RoleId:   request.RoleId,
	}

	db := database.DB
	if err := db.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": false, "message": err, "data": nil})
	}

	user.Password = ""

	return c.JSON(fiber.Map{"status": true, "message": "success", "data": user})
}

func UserEdit(c *fiber.Ctx, request *request.UserUpdate) error {
	id := c.Params("id")

	if !helper.ValidEmail(request.Email) {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": "Email not valid !", "data": nil})
	}

	db := database.DB

	var user model.User
	if err := db.First(&user, "id = ?", id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": false, "message": "User not found !", "data": nil})
	}

	data := model.User{
		Name:     request.Name,
		Username: request.Username,
		Email:    request.Email,
		RoleId:   request.RoleId,
	}

	if err := db.Model(&user).Updates(&data).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": false, "message": err, "data": nil})
	}

	user.Password = ""

	return c.JSON(fiber.Map{"status": true, "message": "success", "data": user})
}

func UserFind(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DB

	var user model.User
	if err := db.First(&user, "id = ?", id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": false, "message": "User not found !", "data": nil})
	}

	var role model.Role
	if err := db.First(&role, "id = ?", user.RoleId).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": false, "message": "Role not found !", "data": nil})
	}

	response := response.UserRoleResponse(&user, &role)

	return c.JSON(fiber.Map{"status": true, "message": "success", "data": response})
}

func UserDelete(c *fiber.Ctx) error {
	id := c.Params("id")

	var user model.User

	db := database.DB
	if err := db.First(&user, "id = ?", id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": false, "message": "User not found !", "data": nil})
	}

	if err := db.Delete(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": false, "message": err, "data": nil})
	}

	return c.JSON(fiber.Map{"status": true, "message": "success", "data": nil})
}
