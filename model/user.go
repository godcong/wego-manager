package model

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego-auth-manager/util"
	log "github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
)

// UserTypeAdmin ...
const UserTypeAdmin = "admin"

// UserTypeUser ...
const UserTypeUser = "user"

// Login ...
type Login struct {
	Username string `json:"username" ` //用户名
	Password string `json:"password" ` //密码
}

// User ...
type User struct {
	Model       `xorm:"extends" json:",inline"`
	Block       bool   `json:"block" xorm:"notnull default(false) comment(禁止访问)"`    //禁止访问
	UserType    string `json:"user_type" xorm:"user_type"`                           //用户类型
	Nickname    string `json:"nickname" xorm:"nickname"`                             //名称
	Username    string `json:"username" xorm:"username notnull default('') unique"`  //用户名
	Email       string `json:"email" xorm:"email notnull default('')"`               //邮件
	Mobile      string `json:"mobile" xorm:"mobile notnull default('')"`             //移动电话
	Password    string `json:"password" xorm:"password"`                             //密码
	Certificate string `json:"certificate" xorm:"certificate"`                       //证书
	PrivateKey  string `json:"private_key" xorm:"private_key"`                       //私钥
	LoginIP     string `json:"login_ip" xorm:"login_ip"`                             //本次登录IP
	Sign        string `json:"sign" xorm:"notnull unique default('')  comment(分享码)"` //分享码
	Token       string `json:"-" xorm:"varchar(1024) token"`                         //Token
	Salt        string `json:"-" xorm:"slat"`                                        //盐值
}

// NewUser ...
func NewUser(id string) *User {
	return &User{Model: Model{
		ID: id,
	}}
}

// Get ...
func (obj *User) Get() (bool, error) {
	return Get(nil, obj)
}

// Update ...
func (obj *User) Update(cols ...string) (int64, error) {
	return Update(nil, obj.ID, obj)
}

// Users ...
func (obj *User) Users() ([]*User, error) {
	var users []*User
	err := DB().Table(obj).Find(&users)
	if err != nil {
		return nil, xerrors.Errorf("find: %w", err)
	}
	return users, nil
}

// Permissions ...
func (obj *User) Permissions() ([]*Permission, error) {
	var permissions []*Permission
	session := DB().Table(&Permission{}).Select("permission.*").
		Join("left", &PermissionUser{}, "permission_user.permission_id = permission.id")

	if obj.ID != "" {
		session = session.Where("permission_user.user_id = ? ", obj.ID)
	}

	err := session.Find(&permissions)
	if err != nil {
		return nil, xerrors.Errorf("relate: %w", err)
	}

	return permissions, nil
}

// CheckPermission ...
func (obj *User) CheckPermission(funcName string) bool {
	var permissions []*Permission
	session := DB().Table(&Permission{}).Select("permission.*").
		Join("left", &PermissionUser{}, "permission_user.permission_id = permission.id").
		Where("permission.slug = ?", funcName)
	if obj.ID != "" {
		session = session.Where("permission_user.user_id = ? ", obj.ID)
	}

	i, err := session.FindAndCount(&permissions)
	if err != nil || i <= 0 {
		return false
	}
	return true
}

// Roles ...
func (obj *User) Roles() ([]*Role, error) {
	var roles []*Role
	session := DB().Table(&Role{}).Select("role.*").
		Join("left", &RoleUser{}, "role_user.role_id = role.id")

	if obj.ID != "" {
		session = session.Where("role_user.user_id = ? ", obj.ID)
	}

	err := session.Find(&roles)
	if err != nil {
		return nil, xerrors.Errorf("relate: %w", err)
	}

	return roles, nil
}

// Validate ...
func (obj *User) Validate(u *Login, key string) bool {
	u.Password = util.SHA256(u.Password, key, obj.Salt)
	session := DB().Table(obj).Where("username = ?", u.Username).And("password = ?", u.Password)

	b, err := session.Exist()
	if err != nil || !b {
		return false
	}
	return true
}

// Properties ...
func (obj *User) Properties() ([]*Property, error) {
	var properties []*Property
	err := DB().Where("property.user_id = ?", obj.ID).Find(&properties)
	if err != nil {
		return nil, xerrors.Errorf("find user properties error : %w", err)
	}
	return properties, nil
}

// UserActivities ...
func (obj *User) UserActivities() ([]*UserActivity, error) {
	var activities []*UserActivity
	err := DB().Where("user_activity.user_id = ?", obj.ID).Find(&activities)
	if err != nil {
		return nil, xerrors.Errorf("find user properties error : %w", err)
	}
	return activities, nil
}

// MustUser ...
func MustUser(user interface{}, b bool) *User {
	if b {
		if v0, b := user.(*User); b {
			log.Printf("%+v\n", v0)
			return v0
		}
	}
	return &User{}
}

// GetUser ...
func GetUser(ctx *gin.Context) *User {
	if v, b := ctx.Get("user"); b {
		if v0, b := v.(*User); b {
			log.Printf("%+v\n", v0)
			return v0
		}
	}
	return nil
}
