package model

import (
	"github.com/godcong/wego-auth-manager/util"
	log "github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"net/url"
)

// UserType ...
type UserType string

// UserTypeAdmin ...
const UserTypeAdmin UserType = "admin"

// UserTypeUser ...
const UserTypeUser UserType = "user"

// Login ...
type Login struct {
	Username string `json:"username" ` //用户名
	Password string `json:"password" ` //密码
}

// User ...
type User struct {
	Model        `xorm:"extends" json:",inline"`
	WechatUserID string   `json:"wechat_user_id" xorm:"wechat_user_id comment(微信用户ID)"`
	Enable       bool     `json:"enable" xorm:"enable"`                                //是否启用
	UserType     UserType `json:"user_type" xorm:"user_type"`                          //用户类型
	Nickname     string   `json:"nickname" xorm:"nickname"`                            //名称
	Username     string   `json:"username" xorm:"username notnull default('') unique"` //用户名
	Email        string   `json:"email" xorm:"email notnull default('')"`              //邮件
	Mobile       string   `json:"mobile" xorm:"mobile notnull default('')"`            //移动电话
	Password     string   `json:"password" xorm:"password"`                            //密码
	Certificate  string   `json:"certificate" xorm:"certificate"`                      //证书
	PrivateKey   string   `json:"private_key" xorm:"private_key"`                      //私钥
	LoginIP      string   `json:"login_ip" xorm:"login_ip"`                            //本次登录IP
	Sign         string   `json:"sign" xorm:"notnull unique default('')  comment(分享码)"`
	Token        string   `json:"-" xorm:"varchar(1024) token"` //Token
	Salt         string   `json:"-" xorm:"slat"`                //盐值
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

// Paginate ...
func (obj *User) Paginate(v url.Values) (*Paginate, error) {
	return &Paginate{}, nil
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

// Property ...
func (obj *User) Property() (*Property, error) {
	var property Property
	b, err := DB().Where("user_property.user_id = ?", obj.ID).Get(&property)
	if err != nil {
		return nil, xerrors.Errorf("find user property error : %w", err)
	}
	if !b {
		return nil, xerrors.New("find user property null")
	}
	return &property, nil
}

// MustUser ...
func MustUser(user interface{}, b bool) *User {
	if b {
		if v0, b := user.(*User); b {
			log.Printf("%+v\n", v0)
			return v0
		}
	}
	return nil
}
