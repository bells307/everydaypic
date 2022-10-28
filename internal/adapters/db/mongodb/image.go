package mongodb

import (
	"context"
	"fmt"
	"log"

	"github.com/bells307/everydaypic/internal/domain/dto"
	"github.com/bells307/everydaypic/internal/domain/entity"
	"github.com/bells307/everydaypic/internal/domain/usecase"
	"github.com/bells307/everydaypic/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type imageStorage struct {
	db *mongodb.MongoDB
}

func NewImageStorage(db *mongodb.MongoDB) *imageStorage {
	return &imageStorage{db}
}

func (s *imageStorage) GetImages(ctx context.Context, dto dto.GetImages) ([]entity.Image, error) {
	filter := bson.M{}

	// Формируем фильтр по ID
	ids := dto.ID
	oids := bson.A{}
	for i := range ids {
		oid, err := primitive.ObjectIDFromHex(ids[i])
		if err != nil {
			return []entity.Image{}, fmt.Errorf("can't convert %s to ObjectID: %v", ids[i], err)
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
		filter["filename"] = bson.M{"$in": fileNames}
	}

	// TODO: убрать хардкод ...
	cur, err := s.db.Find(ctx, "fs.files", filter)
	if err != nil {
		return []entity.Image{}, fmt.Errorf("can't find image: %v", err)
	}
	defer cur.Close(ctx)

	var imgs []entity.Image
	if err := cur.All(ctx, &imgs); err != nil {
		return []entity.Image{}, fmt.Errorf("error decoding mongodb cursor: %v", err)
	}

	if len(imgs) == 0 {
		return []entity.Image{}, usecase.ErrNotFound
	}

	return imgs, nil
}

func (s *imageStorage) CreateImage(ctx context.Context, dto dto.CreateImage) (string, error) {
	log.Println("creating image in mongo ...")

	meta := map[string]any{"name": dto.Name}
	oid, err := s.db.UploadFile(dto.FileName, meta, dto.Data)
	if err != nil {
		return "", fmt.Errorf("error uploading image to mongo bucket: %v", err)
	}
	return oid.Hex(), nil
}

func (s *imageStorage) DeleteImage(ctx context.Context, id string) error {
	log.Printf("deleting image %s in mongo ...\n", id)
	return nil
}

func (s *imageStorage) DownloadImage(ctx context.Context, id string) ([]byte, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return []byte{}, fmt.Errorf("can't convert id %v to hex string", id)
	}

	data, err := s.db.DownloadFile(ctx, oid)
	if err == mongo.ErrNoDocuments {
		return []byte{}, usecase.ErrNotFound
	}
	return data, nil
}
