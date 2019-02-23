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
	City           string   `xorm:"notnull default('0') city" json:"city"`
	Country        string   `xorm:"notnull default('0') country" json:"country"`
	HeadImgURL     string   `xorm:"notnull default('0') head_img_url" json:"headimgurl"`
	Language       string   `xorm:"notnull default('0') language" json:"language"`
	Nickname       string   `xorm:"notnull default('0') nickname" json:"nickname"`
	OpenID         string   `xorm:"notnull default('0') unique open_id" json:"openid"`
	Privilege      []string `xorm:"notnull privilege" json:"privilege"`
	Province       string   `xorm:"notnull default('0') province" json:"province"`
	Sex            uint     `xorm:"notnull default(0) sex" json:"sex"`
	Subscribe      int      `xorm:"notnull default(0) subscribe" json:"subscribe"`
	SubscribeTime  uint32   `xorm:"notnull default(0) subscribe_time" json:"subscribe_time"`
	UnionID        string   `xorm:"notnull default('0') union_id" json:"unionid"`
	Remark         string   `xorm:"notnull default('0') remark" json:"remark"`
	GroupID        int      `xorm:"notnull default(0) group_id" json:"groupid"`
	TagIDList      []int    `xorm:"notnull tag_id_list" json:"tagid_list"`
	SubscribeScene string   `xorm:"notnull default('') subscribe_scene" json:"subscribe_scene"`
	QrScene        int      `xorm:"notnull default(0) qr_scene" json:"qr_scene"`
	QrSceneStr     string   `xorm:"notnull default('0') qr_scene_str" json:"qr_scene_str"`
}

// WechatUser ...
type WechatUser struct {
	Model           `xorm:"extends" json:",inline"`
	AppID           string `xorm:"notnull default('') comment(appid) app_id" json:"appid,omitempty"`                  //appid
	WechatType      string `xorm:"notnull default('') comment(微信or小程序用户标识) wechat_type" json:"wechat_type,omitempty"` //WechatType
	*WechatUserInfo `xorm:"extends" json:",inline"`
}

// Get ...
func (obj *WechatUser) Get() (bool, error) {
	return Get(nil, obj)
}

// UserFromHook ...
func UserFromHook(info *WechatUserInfo, id string, wtype string) *WechatUser {
	return &WechatUser{
		AppID:          id,
		WechatType:     wtype,
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
