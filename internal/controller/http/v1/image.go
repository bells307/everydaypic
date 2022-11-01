package v1

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/bells307/everydaypic/internal/dto"
	"github.com/bells307/everydaypic/internal/entity"
	"github.com/bells307/everydaypic/internal/service"
	"github.com/gin-gonic/gin"
)

// Контроллер изображений
type imageHandler struct {
	imageService ImageService
}

// Интерфейс сервиса работы с изображениями
type ImageService interface {
	// Получить изображения
	GetImages(ctx context.Context, dto dto.GetImages) ([]entity.Image, error)
	// Создать изображение
	CreateImage(ctx context.Context, dto dto.CreateImage) (entity.Image, error)
	// Удалить изображение
	DeleteImage(ctx context.Context, id string) error
	// Скачать изображение
	DownloadImage(ctx context.Context, id string) ([]byte, error)
}

func NewImageHandler(imageService ImageService) *imageHandler {
	return &imageHandler{imageService}
}

// Регистрация роутов контроллера
func (h *imageHandler) Register(e *gin.Engine) {
	v1 := e.Group("/v1")
	{
		images := v1.Group("/image")
		{
			images.GET("", h.getImages)
			images.GET("/:id/download", h.downloadImage)
			images.POST("", h.createImage)
			images.DELETE("/:id", h.deleteImage)
		}
	}

}

// Получить изображения
func (h *imageHandler) getImages(c *gin.Context) {
	var getImages dto.GetImages

	if err := c.Bind(&getImages); err != nil {
		return
	}

	imgs, err := h.imageService.GetImages(c.Request.Context(), getImages)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrNotFound):
			c.AbortWithStatus(http.StatusNotFound)
			return
		default:
			fmt.Printf("err: %v\n", err)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}
	c.JSON(http.StatusOK, imgs)
}

// Создать изображение
func (h *imageHandler) createImage(c *gin.Context) {
	var createImage CreateImage

	if err := c.ShouldBind(&createImage); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	file := createImage.File
	src, err := file.Open()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer src.Close()

	var data []byte
	data, err = io.ReadAll(src)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	dto := dto.CreateImage{
		Name:     createImage.Name,
		FileName: createImage.Filename,
		Data:     data,
	}

	img, err := h.imageService.CreateImage(c.Request.Context(), dto)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, img)
}

// Удалить изображение
func (h *imageHandler) deleteImage(c *gin.Context) {
	id := c.Param("id")
	err := h.imageService.DeleteImage(c.Request.Context(), id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}

// Скачать изображение
func (h *imageHandler) downloadImage(c *gin.Context) {
	id := c.Param("id")
	data, err := h.imageService.DownloadImage(c.Request.Context(), id)

	if err != nil {
		switch {
		case errors.Is(err, service.ErrNotFound):
			c.AbortWithStatus(http.StatusNotFound)
			return
		default:
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}
	c.Data(http.StatusOK, "application/octet-stream", data)
}
