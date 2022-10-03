package response

import (
	"api-fiber-gorm/helper"
	"api-fiber-gorm/model"
	"time"
)

type User struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	RoleId    string    `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Role struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Permission struct {
	Id        string `json:"id"`
	ParenMenu string `json:"parent_menu"`
	ParentId  string `json:"parent_id"`
	Name      string `json:"name"`
	Alias     string `json:"alias"`
	Url       string `json:"url"`
	Icon      string `json:"icon"`
}

type RolePermission struct {
	Id         string       `json:"id"`
	Name       string       `json:"name"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
	Permission []Permission `json:"permissions"`
}

func UserRoleResponse(user *model.User, role *model.Role) interface{} {
	data := map[string]interface{}{
		"id":         user.Id,
		"name":       user.Name,
		"username":   user.Username,
		"email":      user.Email,
		"role_id":    user.RoleId,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
		"role":       role,
	}

	return data
}

func UserRolePermissionResponse(user *model.User, role *model.Role, permissions []model.Permission) interface{} {
	actions := helper.UserActions(permissions)

	rolePermission := map[string]interface{}{
		"id":          role.Id,
		"name":        role.Name,
		"created_at":  role.CreatedAt,
		"updated_at":  role.UpdatedAt,
		"permissions": permissions,
	}

	data := map[string]interface{}{
		"id":         user.Id,
		"name":       user.Name,
		"username":   user.Username,
		"email":      user.Email,
		"role_id":    user.RoleId,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
		"role":       rolePermission,
		"actions":    actions,
	}

	return data
}
