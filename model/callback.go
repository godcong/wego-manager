package model

// Callback ...
type Callback struct {
	Model    `xorm:"extends"`
	Type     string
	BackType string
	Address  string
	Port     string
}

// NewCallback ...
func NewCallback(id string) *Callback {
	return &Callback{
		Model: Model{
			ID: id,
		},
	}
}
