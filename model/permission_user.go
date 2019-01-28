package model

// PermissionUser ...
type PermissionUser struct {
	Model        `xorm:"-"`
	PermissionID string `json:"permission_id" xorm:"permission_id notnull uuid"`
	UserID       string `json:"user_id" xorm:"user_id notnull uuid"`
}

// Relate ...
func (obj *PermissionUser) Relate(permission *Permission, user *User) error {
	return DB().Table(permission).
		Join("left", obj, "permission_user.permission_id = permission.id").
		Join("left", user, "permission_user.user_id = user.id").
		Find(permission, user)
}
