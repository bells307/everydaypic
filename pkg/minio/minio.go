package minio

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinIOClient struct {
	client *minio.Client
}

func NewMinIOClient(endpoint, accessKeyID, secretAccessKey string) (*MinIOClient, error) {
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create minio client: %v", err)
	}

	return &MinIOClient{client}, nil
}

func (c *MinIOClient) UploadFile(ctx context.Context, objectName, fileName, bucketName string, fileSize int64, reader io.ReadSeeker) error {
	exists, errBucketExists := c.client.BucketExists(ctx, bucketName)
	if errBucketExists != nil || !exists {
		err := c.client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return fmt.Errorf("failed to create new bucket: %w", err)
		}
	}

	contentType, err := detectContentType(reader)
	if err != nil {
		return fmt.Errorf("failed to detect content type: %w", err)
	}

	_, err = c.client.PutObject(ctx, bucketName, objectName, reader, fileSize,
		minio.PutObjectOptions{
			UserMetadata: map[string]string{
				"name": fileName,
			},
			ContentType: contentType,
		},
	)

	if err != nil {
		return fmt.Errorf("failed to upload file: %w", err)
	}

	return nil
}

func (c *MinIOClient) GetFileURL(ctx context.Context, bucketName, objectName string, expires time.Duration) (*url.URL, error) {
	return c.client.PresignedGetObject(ctx, bucketName, objectName, expires, url.Values{})
}

func detectContentType(reader io.ReadSeeker) (string, error) {
	dst := bytes.NewBuffer([]byte{})

	if _, err := io.CopyN(dst, reader, 512); err != nil {
		return "", err
	}

	if _, err := reader.Seek(0, 0); err != nil {
		return "", err
	}

	return http.DetectContentType(dst.Bytes()), nil
}
