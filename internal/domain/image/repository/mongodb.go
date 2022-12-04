package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/bells307/everydaypic/internal/domain/image/dto"
	"github.com/bells307/everydaypic/internal/domain/image/model"
	"github.com/bells307/everydaypic/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const COLLECTION_NAME = "image"

// Хранилище
type imageMongoDBRepository struct {
	mongoDBClient *mongodb.MongoDBClient
}

func NewImageMongoDBRepository(mongoDBClient *mongodb.MongoDBClient) *imageMongoDBRepository {
	return &imageMongoDBRepository{mongoDBClient}
}

// Добавить изображение
func (r *imageMongoDBRepository) Add(ctx context.Context, name, fileName, userID string) (model.Image, error) {
	created := time.Now()

	doc := bson.M{
		"name":     name,
		"fileName": fileName,
		"userID":   userID,
		"created":  created,
	}

	res, err := r.mongoDBClient.InsertOne(ctx, COLLECTION_NAME, doc)
	if err != nil {
		return model.Image{}, fmt.Errorf("error inserting image to mongodb collection: %v", err)
	}

	return model.Image{
		ID:       res.InsertedID.(primitive.ObjectID).Hex(),
		Name:     name,
		FileName: fileName,
		UserID:   userID,
		Created:  created,
	}, nil
}

// Удалить изображение
func (r *imageMongoDBRepository) Delete(ctx context.Context, imageID string) error {
	panic("not implemented")
}

// Получить изображения по фильтру
func (r *imageMongoDBRepository) Get(ctx context.Context, dto dto.GetImages) ([]model.Image, error) {
	filter := bson.M{}

	// Формируем фильтр по ID
	ids := dto.ID
	oids := bson.A{}
	for i := range ids {
		oid, err := primitive.ObjectIDFromHex(ids[i])
		if err != nil {
			return []model.Image{}, fmt.Errorf("can't convert %s to ObjectID: %v", ids[i], err)
		} else {
			oids = append(oids, oid)
		}
	}

	if len(oids) > 0 {
		filter["_id"] = bson.M{"$in": oids}
	}

	// Формируем фильтр по именам файлов
	fileNames := bson.A{}
	for i := range dto.FileName {
		fileNames = append(fileNames, dto.FileName[i])
	}

	if len(fileNames) > 0 {
		filter["fileName"] = bson.M{"$in": fileNames}
	}

	cur, err := r.mongoDBClient.Find(ctx, COLLECTION_NAME, filter)
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
	imageOID, err := primitive.ObjectIDFromHex(imageID)
	if err != nil {
		return false, fmt.Errorf("can't convert %s to ObjectID: %v", imageID, err)
	}

	c, err := r.mongoDBClient.GetCount(ctx, COLLECTION_NAME, bson.M{"_id": imageOID})
	if err != nil {
		return false, fmt.Errorf("error counting collection %s with ID %s: %s", COLLECTION_NAME, imageID, err)
	}

	return c > 0, nil
}
