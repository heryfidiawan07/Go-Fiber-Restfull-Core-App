package helper

import (
	"api-fiber-gorm/database"
	"api-fiber-gorm/model"

	"github.com/gofiber/fiber/v2"
)

func Can(c *fiber.Ctx, action string) (bool, string) {
	tokenUserId, err := JwtParse(c)
	if len(err) > 1 {
		return false, "Error JwtParse"
	}

	db := database.DB

	var user model.User
	if err := db.First(&user, "id = ?", tokenUserId).Error; err != nil {
		// return c.Status(404).JSON(fiber.Map{"status": false, "message": "User not found !", "data": nil})
		return false, "User not found !"
	}

	permissions, err := Permissions(user.RoleId)
	if len(err) > 1 {
		// return c.Status(404).JSON(fiber.Map{"status": false, "message": err, "data": nil})
		return false, err
	}

	if action == "me" || action == "except" {
		return true, ""
	}

	for _, value := range permissions {
		if action == value.Name {
			return true, ""
		}
	}

	// return c.Status(403).JSON(fiber.Map{"status": false, "message": "Permission denied !", "data": nil})
	return false, "Permission denied !"
}

func Permissions(roleId string) ([]model.Permission, string) {
	db := database.DB

	var rolePermissions []model.RolePermission
	if err := db.Where("role_id = ?", roleId).Find(&rolePermissions).Error; err != nil {
		return nil, "Role Permission not found !"
	}

	rolePermissionId := make([]string, len(rolePermissions))
	for key, value := range rolePermissions {
		rolePermissionId[key] = value.PermissionId
	}

	var permissions []model.Permission
	if err := db.Where("id IN ?", rolePermissionId).Find(&permissions).Error; err != nil {
		return nil, "Permission not found !"
	}

	return permissions, ""
}

func UserActions(permissions []model.Permission) []string {
	actions := make([]string, len(permissions))
	for key, value := range permissions {
		actions[key] = value.Name
	}
	return actions
}
