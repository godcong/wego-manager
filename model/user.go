package model

import (
	"golang.org/x/exp/xerrors"
	"net/url"
)

// User ...
type User struct {
	Model         `xorm:"extends" json:",inline"`
	Block         bool   `xorm:"block"`           //禁止访问
	Nickname      string `xorm:"nickname"`        //名称
	Username      string `xorm:"username"`        //用户名
	Email         string `xorm:"email"`           //邮件
	Mobile        string `xorm:"mobile"`          //移动电话
	IDCardFacade  string `xorm:"id_card_facade"`  //身份证(正)
	IDCardObverse string `xorm:"id_card_obverse"` //身份证(反)
	Password      string `xorm:"password"`        //密码
	Certificate   string `xorm:"certificate"`     //证书
	PrivateKey    string `xorm:"private_key"`     //私钥
	LoginIP       string `xorm:"login_ip"`        //本次登录IP
	Token         string `xorm:"token"`
}

// NewUser ...
func NewUser(id string) *User {
	return &User{Model: Model{
		ID: id,
	}}
}

// Paginate ...
func (obj *User) Paginate(v url.Values) (*Paginate, error) {
	return &Paginate{}, nil
}

// Users ...
func (obj *User) Users() ([]*User, error) {
	users := new([]*User)
	err := DB().Table(obj).Find(users)
	if err != nil {
		return nil, xerrors.Errorf("find: %w", err)
	}
	return *users, nil
}

// Permissions ...
func (obj *User) Permissions() ([]*Permission, error) {
	var permissions []*Permission
	session := DB().Table(&Permission{}).Select("permission.*").
		Join("left", &PermissionUser{}, "permission_user.permission_id = permission.id").
		Join("left", obj, "permission_user.user_id = user.id")

	if obj.ID != "" {
		session = session.Where("user.id = ? ", obj.ID)
	}

	err := session.Find(&permissions)
	if err != nil {
		return nil, xerrors.Errorf("relate: %w", err)
	}

	return permissions, nil
}

// Roles ...
func (obj *User) Roles() ([]*Role, error) {
	var roles []*Role
	session := DB().Table(&Role{}).Select("role.*").
		Join("left", &RoleUser{}, "role_user.role_id = role.id").
		Join("left", obj, "role_user.user_id = user.id")

	if obj.ID != "" {
		session = session.Where("user.id = ? ", obj.ID)
	}

	err := session.Find(&roles)
	if err != nil {
		return nil, xerrors.Errorf("relate: %w", err)
	}

	return roles, nil
}
