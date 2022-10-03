package model

type RolePermissionCombine struct {
	Role        Role         `json:"role"`
	Permissions []Permission `json:"permissions"`
}
