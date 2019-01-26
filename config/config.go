package config

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"net/url"
)

var globalConfig *Configure

// Config ...
func Config() *Configure {
	return globalConfig
}

// InitConfig ...
func InitConfig(path string) *Configure {
	globalConfig = initLoader(path)
	return globalConfig
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
	Username string `json:"username"`
	Password string `json:"password"`
	Schema   string `json:"schema"`
	Location string `json:"location"`
}

// Source ...
func (d *Database) Source() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?loc=%s&charset=utf8&parseTime=true",
		d.Username, d.Password, d.Addr, d.Port, d.Schema, d.Location)
}

// SourceForCreate ...
func (d *Database) SourceForCreate() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/?loc=%s&charset=utf8&parseTime=true",
		d.Username, d.Password, d.Addr, d.Port, d.Location)
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
	return &Configure{
		Database: Database{
			Type:     "mysql",
			Addr:     "localhost",
			Port:     "3306",
			Username: "root",
			Password: "111111",
			Schema:   "auth",
			Location: url.QueryEscape("Asia/Shanghai"),
		},
		REST: REST{},
	}
}

// MustString ...
func MustString(v, def string) string {
	if v == "" {
		return def
	}
	return v
}
