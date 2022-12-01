package repository

import "github.com/bells307/everydaypic/pkg/minio"

type MinIOImageStorage struct {
	client *minio.MinIOClient
}
