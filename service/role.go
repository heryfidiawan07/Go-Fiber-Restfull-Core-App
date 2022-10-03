package service

import (
	"api-fiber-gorm/database"
	"api-fiber-gorm/helper"
	"api-fiber-gorm/model"
	"api-fiber-gorm/request"
	"api-fiber-gorm/response"

	"github.com/gofiber/fiber/v2"
)

func RoleIndex(c *fiber.Ctx) error {
	var roles []response.Role

	db := database.DB
	if err := db.Find(&roles).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": false, "message": err, "data": nil})
	}

	return c.JSON(fiber.Map{"status": true, "message": nil, "data": roles})
}

func RoleCreate(c *fiber.Ctx, request *request.RoleStore) error {
	role := model.Role{
		Name: request.Name,
	}

	db := database.DB
	if err := db.Create(&role).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": false, "message": err, "data": nil})
	}

	for _, value := range request.Permissions {
		permissions := model.RolePermission{
			RoleId:       role.Id,
			PermissionId: value,
		}
		if err := db.Create(&permissions).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"status": false, "message": err, "data": nil})
		}
	}

	return c.JSON(fiber.Map{"status": true, "message": "success", "data": role})
}

func RoleEdit(c *fiber.Ctx, request *request.RoleUpdate) error {
	id := c.Params("id")

	var role model.Role

	db := database.DB
	if err := db.First(&role, "id = ?", id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": false, "message": "User not found !", "data": nil})
	}

	data := model.Role{
		Name: request.Name,
	}

	if err := db.Model(&role).Updates(&data).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": false, "message": err, "data": nil})
	}

	if err := db.Delete(model.RolePermission{}, "role_id = ?", id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": false, "message": err, "data": nil})
	}

	for _, value := range request.Permissions {
		permissions := model.RolePermission{
			RoleId:       role.Id,
			PermissionId: value,
		}
		if err := db.Create(&permissions).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"status": false, "message": err, "data": nil})
		}
	}

	return c.JSON(fiber.Map{"status": true, "message": "success", "data": role})
}

func RoleFind(c *fiber.Ctx) error {
	id := c.Params("id")

	var role model.Role

	db := database.DB
	if err := db.First(&role, "id = ?", id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": false, "message": "Role not found !", "data": nil})
	}

	permissions, err := helper.Permissions(role.Id)
	if len(err) > 1 {
		return c.Status(404).JSON(fiber.Map{"status": false, "message": err, "data": nil})
	}

	data := map[string]interface{}{
		"id":          role.Id,
		"name":        role.Name,
		"created_at":  role.CreatedAt,
		"updated_at":  role.UpdatedAt,
		"permissions": permissions,
	}

	return c.JSON(fiber.Map{"status": true, "message": "success", "data": data})
}

func RoleDelete(c *fiber.Ctx) error {
	id := c.Params("id")

	var role model.Role

	db := database.DB
	if err := db.First(&role, "id = ?", id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": false, "message": "Role not found !", "data": nil})
	}

	if err := db.Delete(&role).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": false, "message": err, "data": nil})
	}

	if err := db.Delete(model.RolePermission{}, "role_id = ?", id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": false, "message": err, "data": nil})
	}

	return c.JSON(fiber.Map{"status": true, "message": "success", "data": nil})
}
