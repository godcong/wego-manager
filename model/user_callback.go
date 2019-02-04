package model

// BackTypePaid ...
const (
	BackTypePaid    = "paid"
	BackTypeScanned = "scanned"
	BackTypeRefuned = "refund"
)

// UserCallback ...
type UserCallback struct {
	URI      string `xorm:"uri"`
	BackType string `xorm:"back_type"`
}
