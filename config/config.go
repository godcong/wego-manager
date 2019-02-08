package config

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"net/url"
	"strconv"
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
	ShowSQL  bool   `toml:"show_sql"`
	Type     string `toml:"type"`
	Addr     string `toml:"addr"`
	Port     string `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	Schema   string `toml:"schema"`
	Location string `toml:"location"`
	Charset  string `toml:"charset"`
}

// Source ...
func (d *Database) Source() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?loc=%s&charset=%s&parseTime=true",
		d.Username, d.Password, d.Addr, d.Port, d.Schema, d.Location, d.Charset)
}

// General ...
type General struct {
	TokenKey string `toml:"token_key"`
}

// Configure ...
type Configure struct {
	General  General  `toml:"general"`
	Database Database `toml:"database"`
	REST     REST     `toml:"rest"`
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
		General: General{
			TokenKey: "im-godcong-yl",
		},
		Database: Database{
			ShowSQL:  true,
			Type:     "mysql",
			Addr:     "localhost",
			Port:     "3306",
			Username: "root",
			Password: "111111",
			Schema:   "auth",
			Location: url.QueryEscape("Asia/Shanghai"),
			Charset:  "utf8mb4",
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

// MustInt ...
func MustInt(v string, def int) int {
	i, err := strconv.Atoi(v)
	if err == nil {
		return i
	}
	return def
}
