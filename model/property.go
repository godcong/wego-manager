package model

import (
	"github.com/godcong/wego"
	log "github.com/sirupsen/logrus"
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
	Sign        string   `xorm:"notnull unique sign" json:"sign"` //配置唯一识别码
	AppID       string   `xorm:"notnull unique app_id " json:"app_id"`
	Host        string   `xorm:"host" json:"host"`
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

// Get ...
func (obj *Property) Get() (bool, error) {
	return Get(nil, obj)
}

// Properties ...
func (obj *Property) Properties() ([]*Property, error) {
	var properties []*Property
	err := Table(obj).Find(&properties)
	if err != nil {
		return nil, xerrors.Errorf("find: %w", err)
	}
	return properties, nil
}

// Config ...
func (obj *Property) Config() *wego.Config {
	log.Infof("%+v", *obj)
	var config wego.Config
	config.AccessToken = &wego.AccessTokenConfig{
		GrantType: wego.GrantTypeClient,
		AppID:     obj.AppID,
		AppSecret: obj.AppSecret,
	}

	config.OAuth = &wego.OAuthConfig{
		Scopes:      obj.Scopes,
		RedirectURI: obj.RedirectURI,
	}

	config.SafeCert = &wego.SafeCertConfig{
		Cert:   []byte(obj.PemCert),
		Key:    []byte(obj.PemKEY),
		RootCA: []byte(obj.RootCA),
	}

	config.JSSDK = &wego.JSSDKConfig{
		AppID:       obj.AppID,
		MchID:       obj.MchID,
		Key:         obj.MchKey,
		AccessToken: config.AccessToken,
	}

	config.Payment = &wego.PaymentConfig{
		AppID:     obj.AppID,
		AppSecret: obj.AppSecret,
		MchID:     obj.MchID,
		Key:       obj.MchKey,
		SafeCert:  config.SafeCert,
	}

	config.OfficialAccount = &wego.OfficialAccountConfig{
		AppID:       obj.AppID,
		AppSecret:   obj.AppSecret,
		Token:       obj.Token,
		AesKey:      obj.AesKey,
		AccessToken: config.AccessToken,
		OAuth:       config.OAuth,
	}

	return &config
}
