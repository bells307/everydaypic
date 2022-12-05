package main

import (
	"github.com/bells307/everydaypic/cmd/app/config"
	"github.com/bells307/everydaypic/internal/controller/http/v1/image"
	minioRepo "github.com/bells307/everydaypic/internal/domain/file/repository/minio"
	fileService "github.com/bells307/everydaypic/internal/domain/file/service"
	"github.com/bells307/everydaypic/internal/domain/image/repository/mongodb"
	imageService "github.com/bells307/everydaypic/internal/domain/image/service"
	"github.com/bells307/everydaypic/internal/domain/image/usecase"
	"github.com/bells307/everydaypic/pkg/minio"
	mongoClient "github.com/bells307/everydaypic/pkg/mongodb"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title everydaypic API
// @verion 0.1
func main() {
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		panic(err)
	}

	mongo, err := mongoClient.NewMongoDB(cfg.MongoDB)
	if err != nil {
		panic(err)
	}

	minioClient, err := minio.NewMinIOClient(cfg.MinIO)
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	fileRepo := minioRepo.NewMinIOFileRepository(minioClient)
	fileService := fileService.NewFileService(fileRepo)

	imageRepo := mongodb.NewImageMongoDBRepository(mongo)
	imageService := imageService.NewImageService(imageRepo)

	imageUsecase := usecase.NewImageUsecase(imageService, fileService)
	imageHandler := image.NewImageHandler(imageUsecase)
	imageHandler.Register(router)

	router.Run(cfg.Listen)
}
