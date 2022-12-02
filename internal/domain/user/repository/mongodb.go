package repository

import (
	"context"

	"github.com/bells307/everydaypic/internal/domain/user/model"
	"github.com/bells307/everydaypic/pkg/mongodb"
)

const COLLECTION_NAME = "user"

// Хранилище
type userMongoDBRepository struct {
	mongoDBClient *mongodb.MongoDBClient
}

func NewUserMongoDBRepository(mongoDBClient *mongodb.MongoDBClient) *userMongoDBRepository {
	return &userMongoDBRepository{mongoDBClient}
}

func (r *userMongoDBRepository) Add(ctx context.Context, user model.User) error {
	_, err := r.mongoDBClient.InsertOne(ctx, COLLECTION_NAME, user)
	return err
}

func (r *userMongoDBRepository) Delete(ctx context.Context, userID string) error {
	panic("not yet implemented")
}
