package model

// WechatUser ...
type WechatUser struct {
	Model      `xorm:"extends" json:",inline"`
	Block      bool   `xorm:"notnull default(false) comment(禁止访问)"`                           //禁止访问
	AppID      string `xorm:"notnull default('') comment(appid)" json:"appid,omitempty"`      //appid
	OpenID     string `xorm:"notnull default('') comment(用户ID)" json:"openid,omitempty"`      //"openid": "oLVPpjqs9BhvzwPj5A-vTYAX
	UnionID    string `xorm:"notnull default('') comment(唯一ID)" json:"unionid,omitempty"`     //unionid
	Type       int    `xorm:"notnull default(0) comment(微信or小程序用户标识)" json:",omitempty"`      //type
	Nickname   string `xorm:"notnull default('') comment(昵称)" json:"nickname,omitempty"`      //"nickname": "刺猬宝宝",
	Sex        int    `xorm:"notnull default(0) comment(性别)" json:"gender,omitempty"`         //"sex": 1,
	Age        int    `xorm:"notnull default(0) comment(年龄)" json:"age,omitempty"`            //"age"
	Language   string `xorm:"notnull default('') comment(语言)" json:"language,omitempty"`      //"language": "简体中文",
	City       string `xorm:"notnull default('') comment(城市)" json:"city,omitempty"`          //"city": "深圳",
	Province   string `xorm:"notnull default('') comment(省)" json:"province,omitempty"`       //"province": "广东",
	Country    string `xorm:"notnull default('') comment(市)" json:"country,omitempty"`        //"country": "中国",
	HeadImgURL string `xorm:"notnull default('') comment(头像)" json:"avatarurl,omitempty"`     //"headimgurl": "http://wx.qlogo.cn/m
	Privilege  string `xorm:"notnull default('') comment(用户特权信息)" json:"privilege,omitempty"` //"privilege": []
	Height     string `xorm:"notnull default('') comment(身高)" json:"height,omitempty"`        //身高
	Weight     string `xorm:"notnull default('') comment(体重)" json:"weight,omitempty"`        //体重
	Timestamp  int64  `xorm:"notnull default(0) comment(时间戳)" json:"timestamp,omitempty"`     //时间戳
}
