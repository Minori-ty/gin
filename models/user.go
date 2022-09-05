package models

type Address struct {
	State string `json:"state" bson:"state"`
	City  string `json:"city" bson:"city"`
}
type User struct {
	Id       int     `json:"id" bson:"user_id"`
	Username string  `json:"username" bson:"user_name"`
	Age      int     `json:"age" bson:"user_age"`
	Address  Address `json:"address" bson:"user_address"`
}
