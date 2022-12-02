package dto

import (
	"github.com/bells307/everydaypic/internal/domain/user/model"
	"github.com/google/uuid"
)

type CreateUser struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (dto CreateUser) ToModel() model.User {
	return model.User{
		ID:       uuid.NewString(),
		Name:     dto.Name,
		Username: dto.Username,
		Password: dto.Password,
	}
}
