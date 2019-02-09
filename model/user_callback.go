package model

import "golang.org/x/exp/xerrors"

// BackTypePaid ...
const (
	BackTypePaid    = "paid"
	BackTypeScanned = "scanned"
	BackTypeRefuned = "refund"
)

// UserCallback ...
type UserCallback struct {
	Model    `xorm:"extends"`
	UserID   string `json:"user_id" xorm:"user_id"`
	Ver      string `json:"ver" xorm:"ver"`
	BackURL  string `json:"back_url" xorm:"back_url"`
	Sign     string `json:"sign" xorm:"sign notnull unique"`
	URI      string `json:"uri" xorm:"uri"`
	BackType string `json:"back_type" xorm:"back_type"`
}

// Callbacks ...
func (obj *UserCallback) Callbacks() ([]*UserCallback, error) {
	var backs []*UserCallback
	err := DB().Table(obj).Find(&backs)
	if err != nil {
		return nil, xerrors.Errorf("find: %w", err)
	}
	return backs, nil
}

// NewUserCallback ...
func NewUserCallback(id string) *UserCallback {
	return &UserCallback{Model: Model{
		ID: id,
	}}
}
