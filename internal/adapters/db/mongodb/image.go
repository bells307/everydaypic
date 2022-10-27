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

	// ---
	// var purpose string
	// var startDate, endDate time.Time

	// filter := bson.M{}
	// if purpose != "" {
	// 	filter["purpose"] = purpose
	// }
	// if !startDate.IsZero() && !endDate.IsZero() {
	// 	filter["paymentDate"] = bson.M{
	// 		"$gte": startDate,
	// 		"$lt":  endDate,
	// 	}
	// }
	// ---
	// db.feed.find({
	// 	_id: {
	// 		$in: [ObjectId("55880c251df42d0466919268"), ObjectId("55bf528e69b70ae79be35006")]
	// 	}
	// });
	// ---

	cursor, err = collection.Find(ctx, filter)
	res, err := s.db.Find(ctx, "images", dto)
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
