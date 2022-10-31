package main

import (
	v1 "github.com/bells307/everydaypic/internal/controller/http/v1"
	"github.com/bells307/everydaypic/internal/repository/mongodb"
	"github.com/bells307/everydaypic/internal/service"
	"github.com/bells307/everydaypic/pkg/middleware"
	mongodriver "github.com/bells307/everydaypic/pkg/mongodb"
	"github.com/gin-gonic/gin"
)

func main() {
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
	imageHandler := v1.NewImageHandler(imageService)

	router := gin.Default()
	router.Use(middleware.ErrorHandler)
	imageHandler.Register(router)

	router.Run("localhost:8080")
}
