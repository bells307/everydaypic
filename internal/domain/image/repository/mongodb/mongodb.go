package mongodb

import (
	"context"
	"fmt"

	"github.com/bells307/everydaypic/internal/domain/image/model"
	"github.com/bells307/everydaypic/internal/domain/image/repository/dto"
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
func (r *imageMongoDBRepository) GetByFilter(ctx context.Context, filter dto.GetImagesFilter) ([]model.Image, error) {
	m, err := getImagesDTOToBsonMap(filter)
	if err != nil {
		return []model.Image{}, fmt.Errorf("can't convert filter to bson map: %v", err)
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

func getImagesDTOToBsonMap(f dto.GetImagesFilter) (m bson.M, err error) {
	m = bson.M{}

	// Формируем фильтр по ID
	ids := f.ID
	oids := bson.A{}
	for i := range ids {
		oid, err := primitive.ObjectIDFromHex(ids[i])
		if err != nil {
			return m, fmt.Errorf("can't convert %s to ObjectID: %v", ids[i], err)
		} else {
			oids = append(oids, oid)
		}
	}

	if len(oids) > 0 {
		m["_id"] = bson.M{"$in": oids}
	}

	// Формируем фильтр по именам файлов
	fileNames := bson.A{}
	for i := range f.FileName {
		fileNames = append(fileNames, f.FileName[i])
	}

	if len(fileNames) > 0 {
		m["fileName"] = bson.M{"$in": fileNames}
	}

	return m, nil
}
