package config

import (
	"fmt"
	"github.com/pelletier/go-toml"
)

var globalConfig *Configure

// Config ...
func Config() *Configure {
	return globalConfig
}

// InitConfig ...
func InitConfig(path string) {
	globalConfig = initLoader(path)
}

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
	Port     string `json:"port"`
	DB       string `json:"db"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Source ...
func (d *Database) Source() string {
	return fmt.Sprintf("%s:%s@%s:%s/%s?charset=utf8", d.Username, d.Password, d.Addr, d.Port, d.DB)
}

// Configure ...
type Configure struct {
	Database Database `json:"database"`
	REST     REST     `json:"rest"`
}

func initLoader(path string) *Configure {
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
