package model

type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (User) TableName() string { return "users" }

// This what's usually the domain> model> user.go looks like ~
// We defined the user struct to :
// 1- contain the details of users
// 2- be used in our func
// *struct's name must always start with a capital*
