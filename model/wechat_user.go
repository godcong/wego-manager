package model

import (
	"github.com/gin-gonic/gin"
	"log"
)

// WechatTypeH5 ...
const WechatTypeH5 = "h5"

// WechatTypeProgram ...
const WechatTypeProgram = "program"

// WechatUserInfo ...
type WechatUserInfo struct {
	City           string   `json:"city"`
	Country        string   `json:"country"`
	HeadImgURL     string   `json:"headimgurl"`
	Language       string   `json:"language"`
	Nickname       string   `json:"nickname"`
	OpenID         string   `json:"openid"`
	Privilege      []string `json:"privilege"`
	Province       string   `json:"province"`
	Sex            uint     `json:"sex"`
	Subscribe      int      `json:"subscribe"`
	SubscribeTime  uint32   `json:"subscribe_time"`
	UnionID        string   `json:"unionid"`
	Remark         string   `json:"remark"`
	GroupID        int      `json:"groupid"`
	TagIDList      []int    `json:"tagid_list"`
	SubscribeScene string   `json:"subscribe_scene"`
	QrScene        int      `json:"qr_scene"`
	QrSceneStr     string   `json:"qr_scene_str"`
}

// WechatUser ...
type WechatUser struct {
	Model           `xorm:"extends" json:",inline"`
	UserID          string `xorm:"notnull unique default('') user_id"`
	AppID           string `xorm:"notnull default('') comment(appid)" json:"appid,omitempty"`                        //appid
	WechatType      string `xorm:"notnull default(0) comment(微信or小程序用户标识) wechat_type" json:"wechat_type,omitempty"` //WechatType
	*WechatUserInfo `xorm:"extends" json:",inline"`
}

// Get ...
func (obj *WechatUser) Get() (bool, error) {
	return Get(nil, obj)
}

// UserFromHook ...
func UserFromHook(info *WechatUserInfo, id string, typ int) *WechatUser {
	return &WechatUser{
		Model:          Model{},
		UserID:         "",
		AppID:          id,
		WechatType:     "",
		WechatUserInfo: info,
	}
}

// GetWechatUser ...
func GetWechatUser(ctx *gin.Context) *WechatUser {
	if v, b := ctx.Get("user"); b {
		if v0, b := v.(*WechatUser); b {
			log.Printf("%+v\n", v0)
			return v0
		}
	}
	return nil
}
