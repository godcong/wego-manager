package util

import (
	"github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	"time"
)

// WebToken ...
type WebToken struct {
	UID           string `json:"oid"`
	Username      string `json:"username"`
	Nickname      string `json:"nickname"`
	EffectiveTime int64  `json:"effective_time"`
}

// NewWebToken ...
func NewWebToken(uid string) *WebToken {
	return &WebToken{
		UID:           uid,
		EffectiveTime: time.Now().Unix() + 3600*24*7,
	}
}

// ToToken ...
func ToToken(key string, token *WebToken) (string, error) {
	sub, err := jsoniter.Marshal(token)
	if err != nil {
		return "", err
	}
	jwt, err := EncryptJWT([]byte(key), sub)
	return jwt, err
}

// FromToken ...
func FromToken(key, token string) (*WebToken, error) {
	t := WebToken{}
	sub, err := DecryptJWT([]byte(key), token)
	log.Info("sub", sub)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal([]byte(sub), &t)
	if err != nil {
		if err != nil {
			return nil, err
		}
	}
	return &t, nil
}
