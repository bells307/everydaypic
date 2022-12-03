package repository

import (
	"context"
	"io"
	"net/url"

	"github.com/bells307/everydaypic/pkg/minio"
)

const BUCKET_NAME = "images"

type minIOImageStorage struct {
	client *minio.MinIOClient
}

func NewMinIOImageStorage(client *minio.MinIOClient) *minIOImageStorage {
	return &minIOImageStorage{client}
}

func (s *minIOImageStorage) Upload(ctx context.Context, name string, filename string, fileSize int64, data io.Reader) error {
	return s.client.UploadFile(ctx, name, filename, BUCKET_NAME, fileSize, data)
}

func (s *minIOImageStorage) Delete(ctx context.Context, name string) error {
	panic("not yet implemented")
}

func (s *minIOImageStorage) GetUrl(ctx context.Context, imageID string) (url.URL, error) {
	panic("not yet implemented")
}
