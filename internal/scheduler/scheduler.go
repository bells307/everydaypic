package scheduler

import (
	"context"

	"github.com/bells307/everydaypic/internal/dto"
	"github.com/bells307/everydaypic/internal/entity"
)

// Интерфейс сервиса работы с изображениями
type ImageService interface {
	// Получить изображение
	GetImages(ctx context.Context, dto dto.GetImages) ([]entity.Image, error)
	// Установить изображение дня
	SetImageOfTheDay(ctx context.Context, id string) error
}

// Планировщик, устанавливающий изображение дня
type Scheduler struct {
	imageService ImageService
}

func NewScheduler(imageService ImageService) *Scheduler {
	return &Scheduler{imageService: imageService}
}

func (s *Scheduler) Run() {

}
