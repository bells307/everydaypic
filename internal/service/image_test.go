package service_test

import (
	"context"
	"testing"

	"github.com/bells307/everydaypic/internal/repository/mongodb"
	"github.com/bells307/everydaypic/internal/service"
	mongodriver "github.com/bells307/everydaypic/pkg/mongodb"
)

func TestHello(t *testing.T) {
	mongoCfg := mongodriver.MongoDBConfig{
		Uri:    "mongodb://admin:admin@localhost:27017/",
		DbName: "everydaypic",
	}

	mongo, err := mongodriver.NewMongoDB(mongoCfg)
	if err != nil {
		panic(err)
	}

	imageStorage := mongodb.NewImageStorage(mongo)
	imageService := service.NewImageService(imageStorage)

	err = imageService.SetImageOfTheDay(context.TODO(), "635cf2454931a105ef56bb76")

	if err != nil {
		t.Fatal(err)
	}
}
