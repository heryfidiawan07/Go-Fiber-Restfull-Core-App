package model

type RolePermission struct {
	RoleId       string `gorm:"size:36;index;not null" json:"role_id"`
	PermissionId string `gorm:"size:36;index;not null" json:"permission_id"`
}
