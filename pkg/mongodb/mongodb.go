package mongodb

import (
	"bytes"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Конфигурация mongodb
type MongoDBConfig struct {
	Uri    string
	DbName string
}

type MongoDB struct {
	db *mongo.Database
}

func NewMongoDB(cfg MongoDBConfig) (*MongoDB, error) {
	opts := options.Client().ApplyURI(cfg.Uri)
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil, fmt.Errorf("can't connect to mongodb: %v", err)
	}

	db := client.Database(cfg.DbName)
	if db == nil {
		return nil, fmt.Errorf("can't create a handle to mongodb database %s", cfg.DbName)
	}

	return &MongoDB{
		db,
	}, nil
}

// Найти документ в коллекции
func (m *MongoDB) Find(ctx context.Context, collection string, filter any) (*mongo.Cursor, error) {
	col := m.db.Collection(collection)
	return col.Find(ctx, filter)
}

// Добавить элемент в коллекцию
func (m *MongoDB) InsertOne(ctx context.Context, collection string, obj any) (oid primitive.ObjectID, err error) {
	res, err := m.db.Collection(collection).InsertOne(ctx, obj)
	if err != nil {
		return primitive.ObjectID{}, fmt.Errorf("error inserting to collection %s: %v", collection, err)
	}
	oid = res.InsertedID.(primitive.ObjectID)
	return
}

// Добавить или обновить элемент в коллекции
func (m *MongoDB) Upsert(ctx context.Context, collection string, filter any, obj any) (*mongo.UpdateResult, error) {
	opts := options.Update().SetUpsert(true)
	update := bson.M{"$set": obj}

	col := m.db.Collection(collection)
	if col == nil {
		return nil, fmt.Errorf("cant find collection %s", collection)
	}

	return col.UpdateOne(ctx, filter, update, opts)
}

// Загрузить файл в GridFS
func (m *MongoDB) UploadFile(filename string, meta any, data []byte) (primitive.ObjectID, error) {
	bucket, err := gridfs.NewBucket(m.db)
	if err != nil {
		return primitive.ObjectID{}, fmt.Errorf("error creating gridfs bucket: %v", err)
	}

	opts := options.UploadOptions{Metadata: meta}
	uploadStream, err := bucket.OpenUploadStream(filename, &opts)

	if err != nil {
		return primitive.ObjectID{}, fmt.Errorf("error opening upload stream for %s: %v", filename, err)
	}
	defer uploadStream.Close()

	_, err = uploadStream.Write(data)
	if err != nil {
		return primitive.ObjectID{}, fmt.Errorf("error writing to upload stream for %s: %v", filename, err)
	}

	return uploadStream.FileID.(primitive.ObjectID), nil
}

// Скачать файл из GridFS
func (m *MongoDB) DownloadFile(ctx context.Context, oid primitive.ObjectID) ([]byte, error) {
	// TODO: hardcode
	col := m.db.Collection("fs.files")

	var res bson.M
	find := col.FindOne(ctx, bson.M{"_id": oid})
	if err := find.Err(); err == mongo.ErrNoDocuments {
		return []byte{}, err
	}
	err := find.Decode(&res)
	if err != nil {
		return []byte{}, fmt.Errorf("can't decode SingleResult for file %v: %v", oid.Hex(), err)
	}

	bucket, err := gridfs.NewBucket(m.db)
	if err != nil {
		return []byte{}, fmt.Errorf("error creating gridfs bucket: %v", err)
	}

	var buf bytes.Buffer
	_, err = bucket.DownloadToStream(oid, &buf)
	if err != nil {
		return []byte{}, fmt.Errorf("error download file %s: %v", oid.Hex(), err)
	}

	return buf.Bytes(), nil
}
