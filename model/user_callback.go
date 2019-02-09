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
	UserID   string `xorm:"user_id"`
	Ver      string `xorm:"ver"`
	BackURL  string `xorm:"back_url"`
	Sign     string `xorm:"sign notnull unique"`
	URI      string `xorm:"uri"`
	BackType string `xorm:"back_type"`
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
