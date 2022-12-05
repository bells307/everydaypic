package minio

import (
	"context"
	"net/url"
	"time"

	"github.com/bells307/everydaypic/internal/domain/file/repository/dto"
	"github.com/bells307/everydaypic/pkg/minio"
)

// TODO: в конфигурацию
const URL_EXPIRES = time.Hour

type minIOFileRepository struct {
	client *minio.MinIOClient
}

func NewMinIOFileRepository(client *minio.MinIOClient) *minIOFileRepository {
	return &minIOFileRepository{client}
}

func (s *minIOFileRepository) Upload(ctx context.Context, dto dto.UploadFile) error {
	return s.client.UploadFile(ctx, dto.Name, dto.Filename, dto.Bucket, dto.FileSize, dto.Data)
}

func (s *minIOFileRepository) Delete(ctx context.Context, bucket, name string) error {
	return s.client.DeleteFile(ctx, name, bucket)
}

func (s *minIOFileRepository) GetUrl(ctx context.Context, bucket, name string) (*url.URL, error) {
	return s.client.GetFileURL(ctx, bucket, name, URL_EXPIRES)
}
