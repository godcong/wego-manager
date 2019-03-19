package config

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"net/url"
	"strconv"
)

var globalConfig *Configure

// Service ...
type Service struct {
	EnableHTTP bool   `toml:"enable_http"`
	HostPort   string `toml:"host_port"`
	Type       string `toml:"type"`
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
	Initialized bool     `toml:"initialized"`
	WebToken    WebToken `toml:"web_token"`
	Database    Database `toml:"database"`
	Service     Service  `toml:"service"`
}

// IsInitialized ...
func IsInitialized() bool {
	if globalConfig == nil {
		return false
	}
	return globalConfig.Initialized
}

func initLoader(path string) *Configure {
	var cfg Configure
	def := DefaultConfig()
	tree, err := toml.LoadFile(path)
	if err != nil {
		return def
	}
	err = tree.Unmarshal(cfg)
	if err != nil {
		return def
	}
	//url escape
	cfg.Database.Location = url.QueryEscape(cfg.Database.Location)
	return cfg.ParseDefault(def)
}

// ParseDefault ...
func (cfg *Configure) ParseDefault(def *Configure) *Configure {
	cfg.WebToken.Key = MustString(cfg.WebToken.Key, def.WebToken.Key)
	cfg.Database.Type = MustString(cfg.Database.Type, def.Database.Type)
	cfg.Database.Addr = MustString(cfg.Database.Addr, def.Database.Addr)
	cfg.Database.Port = MustString(cfg.Database.Port, def.Database.Port)
	cfg.Database.Username = MustString(cfg.Database.Username, def.Database.Username)
	cfg.Database.Password = MustString(cfg.Database.Password, def.Database.Password)
	cfg.Database.Schema = MustString(cfg.Database.Schema, def.Database.Schema)
	cfg.Database.Location = MustString(cfg.Database.Location, def.Database.Location)
	cfg.Database.Charset = MustString(cfg.Database.Charset, def.Database.Charset)
	cfg.Database.Prefix = MustString(cfg.Database.Prefix, def.Database.Prefix)
	cfg.Service.HostPort = MustString(cfg.Service.HostPort, def.Service.HostPort)
	cfg.Service.Type = MustString(cfg.Service.Type, def.Service.Type)
	return cfg
}

// DefaultConfig ...
func DefaultConfig() *Configure {
	return &Configure{
		WebToken: WebToken{
			Key: "im-godcong-yelion",
		},
		Database: Database{
			ShowSQL:  true,
			UseCache: true,
			Type:     "mysql",
			Addr:     "localhost",
			Port:     "3306",
			Username: "root",
			Password: "111111",
			Schema:   "auth",
			Location: url.QueryEscape("Asia/Shanghai"),
			Charset:  "utf8mb4",
			Prefix:   "",
		},
		Service: Service{
			EnableHTTP: true,
			HostPort:   "8080",
			Type:       "tcp",
		},
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
	if err == nil || i != 0 {
		return i
	}
	return def
}
