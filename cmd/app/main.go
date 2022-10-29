package main

import (
	"github.com/bells307/everydaypic/internal/adapters/db/mongodb"
	v1 "github.com/bells307/everydaypic/internal/controller/http/v1"
	"github.com/bells307/everydaypic/internal/domain/service"
	"github.com/bells307/everydaypic/internal/domain/usecase"
	mongodriver "github.com/bells307/everydaypic/pkg/mongodb"
	"github.com/bells307/everydaypic/pkg/mongodb/middleware"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title everydaypic API
// @description API Server for everydaypic
// @host      localhost:8080
// @BasePath  /api/v1

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
	imageUsecase := usecase.NewImageUsecase(imageService)
	imageHandler := v1.NewImageHandler(imageUsecase)

	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(middleware.ErrorHandler)
	imageHandler.Register(router)

	router.Run("localhost:8080")
}
