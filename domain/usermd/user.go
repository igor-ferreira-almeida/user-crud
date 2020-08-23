package usermd

type User struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

func NewUser(name string, age int, gender string) User {
	return User{
		Name:   name,
		Age:    age,
		Gender: gender,
	}
}
