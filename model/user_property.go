package model

// Payment ...
type Payment struct {
	Sandbox    bool   `toml:"sandbox" json:"sandbox"`
	AppID      string `toml:"app_id" json:"app_id"`
	MchID      string `toml:"mch_id" json:"mch_id"`
	Key        string `toml:"key" json:"key"`
	NotifyURL  string `toml:"notify_url" json:"notify_url"`
	RefundURL  string `toml:"refund_url" json:"refund_url"`
	CertPath   string `toml:"cert_path" json:"cert_path"`
	KeyPath    string `toml:"key_path" json:"key_path"`
	RootCaPath string `toml:"root_ca_path" json:"root_ca_path"`
	PublicKey  string `toml:"public_key" json:"public_key"`
	PrivateKey string `toml:"private_key" json:"private_key"`
}

// OAuth ...
type OAuth struct {
	Scopes      []string `toml:"scopes"`
	RedirectURI string   `toml:"redirect_uri"`
}

// OpenPlatform ...
type OpenPlatform struct {
	AppID  string `toml:"app_id"`
	Secret string `toml:"secret"`
	Token  string `toml:"token"`
	AesKey string `toml:"aes_key"`
}

// OfficialAccount ...
type OfficialAccount struct {
	AppID  string `toml:"app_id"`
	Secret string `toml:"secret"`
	Token  string `toml:"token"`
	AesKey string `toml:"aes_key"`
}

// MiniProgram ...
type MiniProgram struct {
	AppID  string `toml:"app_id"`
	Secret string `toml:"secret"`
	Token  string `toml:"token"`
	AesKey string `toml:"aes_key"`
}

// UserProperty ...
type UserProperty struct {
	UserID string `json:"user_id"`
	Kind   string `json:"kind"`
}
