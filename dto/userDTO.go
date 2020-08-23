package dto

import "github.com/igor-ferreira-almeida/user-crud/domain/usermd"

type UserDTO struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

func (dto UserDTO) ToModel() usermd.User {
	return usermd.User{
		Name: dto.Name,
		Age: dto.Age,
		Gender: dto.Gender,
	}
}

func NewUserDTO(user usermd.User) UserDTO {
	return UserDTO{
		Name: user.Name,
		Age: user.Age,
		Gender: user.Gender,
	}
}
