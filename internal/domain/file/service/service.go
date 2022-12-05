package service

import (
	"context"
	"net/url"

	repoDTO "github.com/bells307/everydaypic/internal/domain/file/repository/dto"
	"github.com/bells307/everydaypic/internal/domain/file/service/dto"
)

// Интерфейс хранилища файлов
type fileRepository interface {
	// Загрузить файл
	Upload(ctx context.Context, dto repoDTO.UploadFile) error
	// Удалить файл
	Delete(ctx context.Context, bucket, name string) error
	// Получить URL файла
	GetUrl(ctx context.Context, bucket, name string) (*url.URL, error)
}

// Сервис работы с загружаемыми файлами
type FileService struct {
	repo fileRepository
}

func NewFileService(repo fileRepository) *FileService {
	return &FileService{repo}
}

func (s *FileService) Upload(ctx context.Context, dto dto.UploadFile) error {
	return s.repo.Upload(ctx, dto.ToRepoDTO())
}

func (s *FileService) Delete(ctx context.Context, bucket, name string) error {
	return s.repo.Delete(ctx, bucket, name)
}

func (s *FileService) GetUrl(ctx context.Context, bucket, name string) (*url.URL, error) {
	return s.repo.GetUrl(ctx, bucket, name)
}
