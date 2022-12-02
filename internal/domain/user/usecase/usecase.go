package usecase

import (
	"context"

	"github.com/bells307/everydaypic/internal/domain/user/dto"
	"github.com/bells307/everydaypic/internal/domain/user/model"
)

// Интерфейс взаимодействия с сервисом пользователей
type userService interface {
	Add(ctx context.Context, dto dto.CreateUser) (model.User, error)
	Delete(ctx context.Context, userID string) error
}

type UserUsecase struct {
	userService userService
}

// Добавить пользователя
func (u *UserUsecase) AddUser(ctx context.Context, dto dto.CreateUser) (model.User, error) {
	return u.userService.Add(ctx, dto)
}

// Удалить пользователя
func (u *UserUsecase) Delete(ctx context.Context, userID string) error {
	return u.userService.Delete(ctx, userID)
}
