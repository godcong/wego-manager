package model

import (
	"golang.org/x/exp/xerrors"
)

// PermissionUser ...
type PermissionUser struct {
	Model        `xorm:"-"`
	PermissionID string `json:"permission_id" xorm:"permission_id notnull uuid"`
	UserID       string `json:"user_id" xorm:"user_id notnull uuid"`
}

// PermissionUserDetail ...
type PermissionUserDetail struct {
	Permission Permission `xorm:"extends"`
	//PermissionUser PermissionUser `xorm:"extends"`
	User User `xorm:"extends"`
}

// Relate ...
func (obj *PermissionUser) Relate(detail *PermissionUserDetail) error {
	pud := new([]PermissionUserDetail)
	err := DB().Table(&detail.Permission).Select("permission.*, user.*").
		Join("left", obj, "permission_user.permission_id = permission.id").
		Join("left", &detail.User, "permission_user.user_id = user.id").
		Where("user.id = ? and permission.id = ?", obj.UserID, obj.PermissionID).Find(pud)
	if err != nil {
		return xerrors.Errorf("relate:%w", err)
	}
	if pud == nil || len(*pud) == 0 {
		return xerrors.New("null permission user detail")
	}
	*detail = (*pud)[0]
	return nil
}
