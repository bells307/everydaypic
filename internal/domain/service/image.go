package service

import "context"

type ImageStorage interface {
}

type ImageService interface {
}

type imageService struct {
	storage ImageStorage
}

func NewImageService(storage ImageStorage) *imageService {
	return &imageService{storage: storage}
}

func (s *imageService) CreateImage(ctx context.Context)
