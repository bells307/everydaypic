package usecase

import (
	"context"

	image_model "github.com/bells307/everydaypic/internal/domain/image/model"
	"github.com/bells307/everydaypic/internal/domain/user/dto"
	"github.com/bells307/everydaypic/internal/domain/user/model"
)

// Интерфейс взаимодействия с сервисом пользователей
type userService interface {
	Add(ctx context.Context, dto dto.CreateUser) (model.User, error)
	Delete(ctx context.Context, userID string) error
}

type imageService interface {
	GetUserImages(ctx context.Context, userID string) ([]image_model.Image, error)
}

type UserUsecase struct {
	userService  userService
	imageService imageService
}

func NewUserUsecase(userService userService, imageService imageService) *UserUsecase {
	return &UserUsecase{userService, imageService}
}

// Добавить пользователя
func (u *UserUsecase) AddUser(ctx context.Context, dto dto.CreateUser) (model.User, error) {
	return u.userService.Add(ctx, dto)
}

// Удалить пользователя
func (u *UserUsecase) Delete(ctx context.Context, userID string) error {
	return u.userService.Delete(ctx, userID)
}
