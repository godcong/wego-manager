package config

import "github.com/pelletier/go-toml"

// REST ...
type REST struct {
	Enable bool   `toml:"enable"`
	Type   string `toml:"type"`
	Path   string `toml:"path"`
	Port   string `toml:"port"`
}

// Database ...
type Database struct {
	Type     string `json:"type"`
	Addr     string `json:"addr"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Configure ...
type Configure struct {
	Database Database `json:"database"`
	REST     REST     `json:"rest"`
}

// InitLoader ...
func InitLoader(path string) *Configure {
	var cfg Configure
	tree, err := toml.LoadFile(path)
	if err != nil {
		return DefaultConfig()
	}
	err = tree.Unmarshal(&cfg)
	if err != nil {
		return DefaultConfig()
	}
	return &cfg
}

// DefaultConfig ...
func DefaultConfig() *Configure {
	return &Configure{}
}

// MustString ...
func MustString(v, def string) string {
	if v == "" {
		return def
	}
	return v
}
