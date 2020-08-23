package dto

import "github.com/igor-ferreira-almeida/user-crud/domain/usermd"

type UserResponseDTO struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

func (dto UserResponseDTO) ToModel() usermd.User {
	return usermd.User{
		ID: 	dto.ID,
		Name:   dto.Name,
		Age:    dto.Age,
		Gender: dto.Gender,
	}
}

func NewUserResponseDTO(user usermd.User) UserResponseDTO {
	return UserResponseDTO{
		ID:     user.ID,
		Name:   user.Name,
		Age:    user.Age,
		Gender: user.Gender,
	}
}
