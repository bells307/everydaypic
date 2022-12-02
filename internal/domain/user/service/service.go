package service

import (
	"context"

	"github.com/bells307/everydaypic/internal/domain/user/dto"
	"github.com/bells307/everydaypic/internal/domain/user/model"
)

// Интерфейс взаимодействия с репозиторием пользователей
type userRepository interface {
	Add(ctx context.Context, user model.User) error
	Delete(ctx context.Context, userID string) error
}

type userService struct {
	userRepository userRepository
}

func NewUserService(userRepository userRepository) *userService {
	return &userService{userRepository}
}

// Добавить пользователя
func (s *userService) Add(ctx context.Context, dto dto.CreateUser) (model.User, error) {
	user := dto.ToModel()
	return user, s.userRepository.Add(ctx, user)
}

// Удалить пользователя
func (s *userService) Delete(ctx context.Context, userID string) error {
	return s.userRepository.Delete(ctx, userID)
}
