package repository

import (
	"context"
	"fmt"

	"github.com/bells307/everydaypic/internal/model"
	mongoClient "github.com/bells307/everydaypic/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const COLLECTION_NAME = "image"

// Хранилище
type imageMongoDBRepository struct {
	mongoDBClient *mongoClient.MongoDBClient
}

func NewImageMongoDBRepository(mongoDBClient *mongoClient.MongoDBClient) *imageMongoDBRepository {
	return &imageMongoDBRepository{mongoDBClient}
}

// Добавить изображение
func (r *imageMongoDBRepository) Create(ctx context.Context, image model.Image) error {
	_, err := r.mongoDBClient.InsertOne(ctx, COLLECTION_NAME, image)
	return err
}

// Удалить изображение
func (r *imageMongoDBRepository) Delete(ctx context.Context, imageID string) error {
	_, err := r.mongoDBClient.DeleteOne(ctx, COLLECTION_NAME, bson.M{"_id": imageID})
	if err != nil {
		return fmt.Errorf("error deleting image from mongodb collection: %v", err)
	}
	return nil
}

// Получить изображения по фильтру
func (r *imageMongoDBRepository) GetByFilter(ctx context.Context, IDs, fileNames []string) ([]model.Image, error) {
	m := bson.M{}

	// Формируем фильтр по ID
	bsonIDs := bson.A{}
	for i := range IDs {
		oid, err := primitive.ObjectIDFromHex(IDs[i])
		if err != nil {
			return []model.Image{}, fmt.Errorf("can't convert %s to ObjectID: %v", IDs[i], err)
		} else {
			bsonIDs = append(bsonIDs, oid)
		}
	}

	if len(bsonIDs) > 0 {
		m["_id"] = bson.M{"$in": bsonIDs}
	}

	// Формируем фильтр по именам файлов
	bsonfileNames := bson.A{}
	for i := range fileNames {
		bsonfileNames = append(bsonfileNames, fileNames[i])
	}

	if len(bsonfileNames) > 0 {
		m["fileName"] = bson.M{"$in": bsonfileNames}
	}

	cur, err := r.mongoDBClient.Find(ctx, COLLECTION_NAME, m)
	if err != nil {
		return []model.Image{}, fmt.Errorf("can't find image: %v", err)
	}
	defer cur.Close(ctx)

	var imgs []model.Image
	if err := cur.All(ctx, &imgs); err != nil {
		return []model.Image{}, fmt.Errorf("error decoding mongodb cursor: %v", err)
	}

	return imgs, nil
}

func (r *imageMongoDBRepository) CheckExists(ctx context.Context, imageID string) (bool, error) {
	c, err := r.mongoDBClient.GetCount(ctx, COLLECTION_NAME, bson.M{"_id": imageID})
	if err != nil {
		return false, fmt.Errorf("error counting collection %s with ID %s: %s", COLLECTION_NAME, imageID, err)
	}

	return c > 0, nil
}
