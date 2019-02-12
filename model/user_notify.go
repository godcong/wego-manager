package model

import "golang.org/x/xerrors"

// BackTypePaid ...
const (
	BackTypePaid    = "paid"
	BackTypeScanned = "scanned"
	BackTypeRefuned = "refund"
)

// UserNotify ...
type UserNotify struct {
	Model    `xorm:"extends"`
	UserID   string `json:"user_id" xorm:"user_id"`
	Sign     string `json:"sign" xorm:"sign notnull unique"`
	Ver      string `json:"ver" xorm:"ver"`
	BackURL  string `json:"back_url" xorm:"back_url"`
	URI      string `json:"uri" xorm:"uri"`
	BackType string `json:"back_type" xorm:"back_type"`
}

// Callbacks ...
func (obj *UserNotify) Callbacks() ([]*UserNotify, error) {
	var backs []*UserNotify
	err := DB().Table(obj).Find(&backs)
	if err != nil {
		return nil, xerrors.Errorf("find: %w", err)
	}
	return backs, nil
}

// NewUserCallback ...
func NewUserCallback(id string) *UserNotify {
	return &UserNotify{Model: Model{
		ID: id,
	}}
}
