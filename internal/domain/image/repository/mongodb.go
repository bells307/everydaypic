package repository

import (
	"context"
	"fmt"

	"github.com/bells307/everydaypic/internal/domain/image/model"
	"github.com/bells307/everydaypic/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

const COLLECTION_NAME = "images"

// Хранилище
type imageMongoDBRepository struct {
	mongoDBClient *mongodb.MongoDBClient
}

func NewImageMongoDBRepository(mongoDBClient *mongodb.MongoDBClient) *imageMongoDBRepository {
	return &imageMongoDBRepository{mongoDBClient}
}

// Добавить изображение
func (r *imageMongoDBRepository) Add(ctx context.Context, image model.Image) error {
	_, err := r.mongoDBClient.InsertOne(ctx, COLLECTION_NAME, image)
	return err
}

// Удалить изображение
func (r *imageMongoDBRepository) Delete(ctx context.Context, imageID string) error {
	panic("not implemented")
}

// Получить изображение по ID
func (r *imageMongoDBRepository) GetByID(ctx context.Context, imageID string) (model.Image, error) {
	var image model.Image
	if err := r.mongoDBClient.FindOne(ctx, COLLECTION_NAME, bson.M{"_id": imageID}).Decode(&image); err != nil {
		return model.Image{}, fmt.Errorf("can't find image %s by ID: %v", imageID, err)
	}

	return image, nil
}
