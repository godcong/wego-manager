package model

// PermissionRole ...
type PermissionRole struct {
	Model        `xorm:"-"`
	PermissionID string `json:"permission_id" xorm:"permission_id notnull uuid"`
	RoleID       string `json:"role_id" xorm:"role_id notnull uuid"`
}
