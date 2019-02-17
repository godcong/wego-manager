package model

import (
	"golang.org/x/xerrors"
)

// OAuth ...
type OAuth struct {
	Scopes      []string `xorm:"oauth.scopes"`
	RedirectURL string   `xorm:"oauth.redirect_url"`
}

// Property ...
type Property struct {
	Model       `xorm:"extends" json:",inline"`
	UserID      string   `xorm:"notnull unique user_id" json:"user_id"`
	Sign        string   `xorm:"notnull unique sign" json:"sign"`
	AppID       string   `xorm:"notnull unique app_id " json:"app_id"`
	MchID       string   `xorm:"mch_id" json:"mch_id"`
	MchKey      string   `xorm:"mch_key" json:"mch_key"`
	PemCert     string   `xorm:"varchar(2048) pem_cert" json:"pem_cert"`
	PemKEY      string   `xorm:"varchar(2048) pem_key" json:"pem_key"`
	RootCA      string   `xorm:"varchar(2048) root_ca" json:"root_ca"`
	NotifyURL   string   `xorm:"notify_url" json:"notify_url"`
	RefundURL   string   `xorm:"refund_url" json:"refund_url"`
	Kind        string   `xorm:"kind" json:"kind"`
	Sandbox     bool     `xorm:"sandbox" json:"sandbox" `
	AppSecret   string   `xorm:"app_secret" json:"app_secret"`
	Token       string   `xorm:"token" json:"token"`
	AesKey      string   `xorm:"aes_key" json:"aes_key"`
	PublicKey   string   `xorm:"public_key" json:"public_key"`
	PrivateKey  string   `xorm:"private_key" json:"private_key"`
	Scopes      []string `xorm:"scopes" json:"scopes"`
	RedirectURI string   `xorm:"redirect_uri" json:"redirect_uri"`
}

// NewProperty ...
func NewProperty(id string) *Property {
	return &Property{Model: Model{
		ID: id,
	}}
}

// Properties ...
func (obj *Property) Properties() ([]*Property, error) {
	var properties []*Property
	err := DB().Table(obj).Find(&properties)
	if err != nil {
		return nil, xerrors.Errorf("find: %w", err)
	}
	return properties, nil
}
