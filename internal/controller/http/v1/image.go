package v1

import (
	"context"
	"errors"
	"io"
	"net/http"

	"github.com/bells307/everydaypic/internal/domain/dto"
	"github.com/bells307/everydaypic/internal/domain/usecase"
	"github.com/gin-gonic/gin"
)

type imageHandler struct {
	imageUsecase ImageUsecase
}

type ImageUsecase interface {
	CreateImage(ctx context.Context, dto dto.CreateImage) (string, error)
	DeleteImage(ctx context.Context, id string) error
	DownloadImage(ctx context.Context, id string) ([]byte, error)
}

func NewImageHandler(imageUsecase ImageUsecase) *imageHandler {
	return &imageHandler{imageUsecase}
}

func (h *imageHandler) Register(e *gin.Engine) {
	v1 := e.Group("/v1")
	{
		images := v1.Group("/image")
		{
			images.GET("/:id/download", h.downloadImage)
			images.POST("/", h.createImage)
			images.DELETE("/:id", h.deleteImage)
		}
	}

}

func (h *imageHandler) createImage(c *gin.Context) {
	var createImage CreateImage

	if err := c.ShouldBind(&createImage); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	file := createImage.File
	src, err := file.Open()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	defer src.Close()

	var data []byte
	data, err = io.ReadAll(src)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	dto := dto.CreateImage{
		Name:     createImage.Name,
		FileName: createImage.Filename,
		Data:     data,
	}

	id, err := h.imageUsecase.CreateImage(c.Request.Context(), dto)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"id": id})
}

func (h *imageHandler) deleteImage(c *gin.Context) {
	id := c.Param("id")
	err := h.imageUsecase.DeleteImage(c.Request.Context(), id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
}

func (h *imageHandler) downloadImage(c *gin.Context) {
	id := c.Param("id")
	data, err := h.imageUsecase.DownloadImage(c.Request.Context(), id)

	if err != nil {
		switch {
		case errors.Is(err, usecase.ErrNotFound):
			c.AbortWithStatusJSON(http.StatusNotFound, err)
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}
	}
	c.Data(http.StatusOK, "application/octet-stream", data)
}
