package main

import (
	"github.com/bells307/everydaypic/internal/controller/http/v1/image"
	"github.com/bells307/everydaypic/internal/domain/image/repository"
	"github.com/bells307/everydaypic/internal/domain/image/service"
	"github.com/bells307/everydaypic/internal/domain/image/usecase"
	"github.com/bells307/everydaypic/pkg/minio"
	"github.com/bells307/everydaypic/pkg/mongodb"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title everydaypic API
// @verion 0.1

func main() {
	// Конфигурация mongodb
	mongoCfg := mongodb.MongoDBConfig{
		Uri:    "mongodb://admin:admin@localhost:27017/",
		DbName: "everydaypic",
	}

	mongo, err := mongodb.NewMongoDB(mongoCfg)
	if err != nil {
		panic(err)
	}

	minioClient, _ := minio.NewMinIOClient("localhost:9000", "user", "mysuperpass")

	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	imageRepo := repository.NewImageMongoDBRepository(mongo)
	imageFS := repository.NewMinIOImageStorage(minioClient)
	imageService := service.NewImageService(imageRepo, imageFS)
	imageUsecase := usecase.NewImageUsecase(imageService)
	imageHandler := image.NewImageHandler(imageUsecase)
	imageHandler.Register(router)

	router.Run("localhost:8080")
}
