package model

// Spread ...
type Spread struct {
	Model        `xorm:"extends"`
	WechatUserID string `json:"wechat_user_id"`
	ParentSign   string `json:"parent_sign comment(上1级)"`
	ParentSign2  string `json:"parent_sign_2 comment(上2级)"`
	ParentSign3  string `json:"parent_sign_3 comment(上3级)"`
	ParentSign4  string `json:"parent_sign_4 comment(上4级)"`
	ParentSign5  string `json:"parent_sign_5 comment(上5级)"`
	ParentSign6  string `json:"parent_sign_6 comment(上6级)"`
	ParentSign7  string `json:"parent_sign_7 comment(上7级)"`
	ParentSign8  string `json:"parent_sign_8 comment(上8级)"`
	ParentSign9  string `json:"parent_sign_9 comment(上9级)"`
}
