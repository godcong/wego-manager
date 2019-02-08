package model

// BackTypePaid ...
const (
	BackTypePaid    = "paid"
	BackTypeScanned = "scanned"
	BackTypeRefuned = "refund"
)

// UserCallback ...
type UserCallback struct {
	Model    `xorm:"extends"`
	Ver      string `xorm:"version"`
	URI      string `xorm:"uri"`
	BackType string `xorm:"back_type"`
}
