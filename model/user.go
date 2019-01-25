package model

type User struct {
	Model `json:",inline"`
	Name  string `json:"name"`
}

func Users() ([]*User, error) {
	return nil, nil
}

func (m *User) Users() []*User {
	return nil
}
