package model

import "net/url"

// User ...
type User struct {
	Model         `xorm:"extends" json:",inline"`
	Block         bool   `xorm:"block"`           //禁止访问
	Name          string `xorm:"name"`            //名称
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

// PaginateUsers ...
func PaginateUsers(val url.Values) (*Paginate, error) {
	var users []*User
	paginate := ParsePaginate(val)
	err := paginate.Engine().Find(&users)
	paginate.Detail = users
	return paginate, err
}

// Users ...
func (m *User) Users() []*User {
	return nil
}
