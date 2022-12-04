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
	Uri    string `mapstructure:"URI"`
	DbName string `mapstructure:"DBNAME"`
}

type MongoDBClient struct {
	db *mongo.Database
}

func NewMongoDB(cfg MongoDBConfig) (*MongoDBClient, error) {
	opts := options.Client().ApplyURI(cfg.Uri)
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil, fmt.Errorf("can't connect to mongodb: %v", err)
	}

	db := client.Database(cfg.DbName)
	if db == nil {
		return nil, fmt.Errorf("can't create a handle to mongodb database %s", cfg.DbName)
	}

	return &MongoDBClient{
		db,
	}, nil
}

// Найти документы в коллекции
func (m *MongoDBClient) Find(ctx context.Context, collection string, filter any) (*mongo.Cursor, error) {
	col := m.db.Collection(collection)
	return col.Find(ctx, filter)
}

// Найти документ в коллекции
func (m *MongoDBClient) FindOne(ctx context.Context, collection string, filter any) *mongo.SingleResult {
	col := m.db.Collection(collection)
	return col.FindOne(ctx, filter)
}

// Добавить элемент в коллекцию
func (m *MongoDBClient) InsertOne(ctx context.Context, collection string, obj any) (*mongo.InsertOneResult, error) {
	return m.db.Collection(collection).InsertOne(ctx, obj)
}

// Добавить или обновить элемент в коллекции
func (m *MongoDBClient) Upsert(ctx context.Context, collection string, filter any, obj any) (*mongo.UpdateResult, error) {
	opts := options.Update().SetUpsert(true)
	update := bson.M{"$set": obj}

	col := m.db.Collection(collection)
	if col == nil {
		return nil, fmt.Errorf("cant find collection %s", collection)
	}

	return col.UpdateOne(ctx, filter, update, opts)
}

// Загрузить файл в GridFS
func (m *MongoDBClient) UploadFile(filename string, meta any, data []byte) (primitive.ObjectID, error) {
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
func (m *MongoDBClient) DownloadFile(ctx context.Context, oid primitive.ObjectID) ([]byte, error) {
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

// Посчитать количество элементов в коллекции, удовлетворяющих условию
func (m *MongoDBClient) GetCount(ctx context.Context, collection string, filter any) (int64, error) {
	return m.db.Collection(collection).CountDocuments(ctx, filter)
}
