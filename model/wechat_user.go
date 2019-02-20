package model

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego/core"
	"log"
)

// WechatTypeH5 ...
const WechatTypeH5 = "h5"

// WechatTypeProgram ...
const WechatTypeProgram = "program"

// WechatUser ...
type WechatUser struct {
	Model                `xorm:"extends" json:",inline"`
	UserID               string `xorm:"notnull unique default('') user_id"`
	AppID                string `xorm:"notnull default('') comment(appid)" json:"appid,omitempty"`                        //appid
	WechatType           string `xorm:"notnull default(0) comment(微信or小程序用户标识) wechat_type" json:"wechat_type,omitempty"` //WechatType
	*core.WechatUserInfo `xorm:"extends" json:",inline"`
}

// UserFromHook ...
func UserFromHook(info *core.WechatUserInfo, id string, typ int) *WechatUser {
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
