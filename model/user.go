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

// Count ...
func (obj *User) Count() (int64, error) {
	return Count(nil, obj)
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
		return nil, xerrors.Opaque(err)
	}
	return users, nil
}
