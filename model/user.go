package model

type User struct {
	Id   int
	Name string
}

func (u *User) UpdateName(name string) {
	u.Name = name
}
