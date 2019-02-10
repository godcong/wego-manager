package model

import "golang.org/x/exp/xerrors"

// OAuth ...
type OAuth struct {
	Scopes      []string `xorm:"oauth.scopes"`
	RedirectURL string   `xorm:"oauth.redirect_url"`
}

// UserProperty ...
type UserProperty struct {
	Model       `xorm:"extends"`
	UserID      string   `xorm:"user_id" json:"user_id"`
	Sign        string   `xorm:"sign" json:"sign"`
	AppID       string   `xorm:"app_id notnull unique" json:"app_id"`
	MchID       string   `xorm:"mch_id" json:"mch_id"`
	MchKey      string   `xorm:"mch_key" json:"mch_key"`
	Cert        string   `xorm:"cert" json:"cert"`
	Key         string   `xorm:"key" json:"key"`
	RootCA      string   `xorm:"root_ca" json:"root_ca"`
	NotifyURL   string   `xorm:"notify_url" json:"notify_url"`
	RefundURL   string   `xorm:"refund_url" json:"refund_url"`
	Kind        string   `xorm:"kind" json:"kind"`
	Sandbox     bool     `xorm:"sandbox" json:"sandbox" json:"sandbox"`
	Secret      string   `xorm:"secret" json:"secret"`
	Token       string   `xorm:"token" json:"token"`
	AesKey      string   `xorm:"aes_key" json:"aes_key"`
	PublicKey   string   `xorm:"public_key" json:"public_key"`
	PrivateKey  string   `xorm:"private_key" json:"private_key"`
	Scopes      []string `xorm:"scopes" json:"scopes"`
	RedirectURL string   `xorm:"redirect_url" json:"redirect_url"`
}

// NewUserProperty ...
func NewUserProperty(id string) *UserProperty {
	return &UserProperty{Model: Model{
		ID: id,
	}}
}

// Properties ...
func (obj *UserProperty) Properties() ([]*UserProperty, error) {
	var properties []*UserProperty
	err := DB().Table(obj).Find(&properties)
	if err != nil {
		return nil, xerrors.Errorf("find: %w", err)
	}
	return properties, nil
}
