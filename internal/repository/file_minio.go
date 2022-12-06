package repository

import (
	"context"
	"io"
	"net/url"
	"time"

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

func (s *minIOFileRepository) Upload(ctx context.Context, name, fileName, bucket string, fileSize int64, data io.ReadSeeker) error {
	return s.client.UploadFile(ctx, name, fileName, bucket, fileSize, data)
}

func (s *minIOFileRepository) Delete(ctx context.Context, bucket, name string) error {
	return s.client.DeleteFile(ctx, name, bucket)
}

func (s *minIOFileRepository) GetUrl(ctx context.Context, bucket, name string) (*url.URL, error) {
	return s.client.GetFileURL(ctx, bucket, name, URL_EXPIRES)
}
