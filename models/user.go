package models

type User struct {
	Id       int
	Username string
}

func (user User) TableName() string {
	return "user"
}
