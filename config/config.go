package config

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"net/url"
	"strconv"
)

var globalConfig *Configure

// REST ...
type REST struct {
	Enable bool   `toml:"enable"`
	Type   string `toml:"type"`
	Path   string `toml:"path"`
	Port   string `toml:"port"`
}

// HTTP ...
type HTTP struct {
	Enable bool   `toml:"enable"`
	Type   string `toml:"type"`
	Path   string `toml:"path"`
	Port   string `toml:"port"`
}

// Database ...
type Database struct {
	ShowSQL  bool   `toml:"show_sql"`
	UseCache bool   `json:"use_cache"`
	Type     string `toml:"type"`
	Addr     string `toml:"addr"`
	Port     string `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	Schema   string `toml:"schema"`
	Location string `toml:"location"`
	Charset  string `toml:"charset"`
	Prefix   string `toml:"prefix"`
}

// Config ...
func Config() *Configure {
	return globalConfig
}

// InitConfig ...
func InitConfig(path string) *Configure {
	globalConfig = initLoader(path)
	return globalConfig
}

// Source ...
func (d *Database) Source() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?loc=%s&charset=%s&parseTime=true",
		d.Username, d.Password, d.Addr, d.Port, d.Schema, d.Location, d.Charset)
}

// WebToken ...
type WebToken struct {
	Key string `toml:"key"`
}

// Configure ...
type Configure struct {
	WebToken WebToken `toml:"web_token"`
	Database Database `toml:"database"`
	REST     REST     `toml:"rest"`
	HTTP     HTTP     `toml:"rest"`
}

func initLoader(path string) (cfg *Configure) {
	cfg = DefaultConfig()
	tree, err := toml.LoadFile(path)
	if err != nil {
		return
	}
	err = tree.Unmarshal(cfg)
	if err != nil {
		return
	}
	return
}

// DefaultConfig ...
func DefaultConfig() *Configure {
	return &Configure{
		WebToken: WebToken{
			Key: "im-godcong-yelion",
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
