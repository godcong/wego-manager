package model

import "golang.org/x/xerrors"

// BackTypePaid ...
const (
	BackTypePaid    = "paid"
	BackTypeScanned = "scanned"
	BackTypeRefuned = "refund"
)

// Notify ...
type Notify struct {
	Model    `xorm:"extends" json:",inline"`
	UserID   string `json:"user_id" xorm:"user_id"`
	Sign     string `json:"sign" xorm:"sign notnull unique"`
	Ver      string `json:"ver" xorm:"ver"`
	BackURL  string `json:"back_url" xorm:"back_url"`
	URI      string `json:"uri" xorm:"uri"`
	BackType string `json:"back_type" xorm:"back_type"`
}

// NewNotify ...
func NewNotify(id string) *Notify {
	return &Notify{Model: Model{
		ID: id,
	}}
}

// Notifies ...
func (obj *Notify) Notifies() ([]*Notify, error) {
	var backs []*Notify
	err := DB().Table(obj).Find(&backs)
	if err != nil {
		return nil, xerrors.Errorf("find: %w", err)
	}
	return backs, nil
}
