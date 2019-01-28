package model

// RoleUser ...
type RoleUser struct {
	Model  `xorm:"-"`
	RoleID string `json:"role_id" xorm:"role_id notnull uuid"`
	UserID string `json:"user_id" xorm:"user_id notnull uuid"`
}
