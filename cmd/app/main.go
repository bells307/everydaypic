package main

import (
	"github.com/bells307/everydaypic/internal/adapters/db/mongodb"
	v1 "github.com/bells307/everydaypic/internal/controller/http/v1"
	"github.com/bells307/everydaypic/internal/domain/service"
	"github.com/bells307/everydaypic/internal/domain/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	imageStorage := mongodb.NewImageStorage(nil)
	imageService := service.NewImageService(imageStorage)
	imageUsecase := usecase.NewImageUsecase(imageService)
	imageHandler := v1.NewImageHandler(imageUsecase)

	router := gin.Default()
	imageHandler.Register(router)

	router.Run("localhost:8080")
}
