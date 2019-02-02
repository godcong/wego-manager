package model

// Payment ...
type Payment struct {
	AppID      string `xorm:"app_id" json:"app_id"`
	MchID      string `xorm:"mch_id" json:"mch_id"`
	Key        string `xorm:"key" json:"key"`
	NotifyURL  string `xorm:"notify_url" json:"notify_url"`
	RefundURL  string `xorm:"refund_url" json:"refund_url"`
	CertPath   string `xorm:"cert_path" json:"cert_path"`
	KeyPath    string `xorm:"key_path" json:"key_path"`
	RootCaPath string `xorm:"root_ca_path" json:"root_ca_path"`
	PublicKey  string `xorm:"public_key" json:"public_key"`
	PrivateKey string `xorm:"private_key" json:"private_key"`
}

// OAuth ...
type OAuth struct {
	Scopes      []string `xorm:"scopes"`
	RedirectURI string   `xorm:"redirect_uri"`
}

// OpenPlatform ...
type OpenPlatform struct {
	AppID  string `xorm:"app_id"`
	Secret string `xorm:"secret"`
	Token  string `xorm:"token"`
	AesKey string `xorm:"aes_key"`
}

// OfficialAccount ...
type OfficialAccount struct {
	AppID  string `xorm:"app_id"`
	Secret string `xorm:"secret"`
	Token  string `xorm:"token"`
	AesKey string `xorm:"aes_key"`
}

// MiniProgram ...
type MiniProgram struct {
	AppID  string `xorm:"app_id"`
	Secret string `xorm:"secret"`
	Token  string `xorm:"token"`
	AesKey string `xorm:"aes_key"`
}

// UserProperty ...
type UserProperty struct {
	Model           `xorm:"extends"`
	UserID          string `xorm:"user_id"`
	Kind            string `xorm:"kind"`
	Sandbox         bool   `xorm:"sandbox" json:"sandbox"`
	OAuth           `xorm:"json oauth"`
	Payment         `xorm:"json"`
	MiniProgram     `xorm:"json"`
	OfficialAccount `xorm:"json"`
	OpenPlatform    `xorm:"json"`
}

// NewUserProperty ...
func NewUserProperty(id string) *UserProperty {
	return &UserProperty{Model: Model{
		ID: id,
	}}
}
