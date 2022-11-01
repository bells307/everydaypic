package mongodb

import (
	"context"
	"fmt"
	"log"

	"github.com/bells307/everydaypic/internal/dto"
	"github.com/bells307/everydaypic/internal/entity"
	"github.com/bells307/everydaypic/internal/service"
	"github.com/bells307/everydaypic/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Реализация хранилища изображений в MongoDB
type imageStorage struct {
	db *mongodb.MongoDB
}

func NewImageStorage(db *mongodb.MongoDB) *imageStorage {
	return &imageStorage{db}
}

// Получить изображения
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
		return []entity.Image{}, service.ErrNotFound
	}

	return imgs, nil
}

// Создать изображение
func (s *imageStorage) CreateImage(ctx context.Context, dto dto.CreateImage) (entity.Image, error) {
	log.Println("creating image in mongo ...")

	meta := map[string]any{
		"name": dto.Name,
	}

	oid, err := s.db.UploadFile(dto.FileName, meta, dto.Data)
	if err != nil {
		return entity.Image{}, fmt.Errorf("error uploading image to mongo bucket: %v", err)
	}
	img := entity.Image{
		ID:       oid.Hex(),
		Filename: dto.FileName,
		Metadata: meta,
	}
	return img, nil
}

// Удалить изображение
func (s *imageStorage) DeleteImage(ctx context.Context, id string) error {
	log.Printf("deleting image %s in mongo ...\n", id)
	return nil
}

// Скачать изображение
func (s *imageStorage) DownloadImage(ctx context.Context, id string) ([]byte, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return []byte{}, fmt.Errorf("can't convert id %v to hex string", id)
	}

	data, err := s.db.DownloadFile(ctx, oid)
	if err == mongo.ErrNoDocuments {
		return []byte{}, service.ErrNotFound
	}
	return data, nil
}

// Создать изображение дня
func (s *imageStorage) CreateImageOfTheDay(ctx context.Context, dayImg entity.DayImage) error {
	oid, err := primitive.ObjectIDFromHex(dayImg.ID)
	if err != nil {
		return fmt.Errorf("can't convert id %v to hex string", dayImg.ID)
	}

	setAt := primitive.NewDateTimeFromTime(dayImg.SetAt)
	obj := bson.M{
		"_id":        oid,
		"filename":   dayImg.Filename,
		"metadata":   dayImg.Metadata,
		"uploadDate": dayImg.UploadDate,
		"setAt":      setAt,
	}
	s.db.InsertOne(ctx, "dayimage", obj)
	return nil
}
