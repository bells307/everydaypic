package v1

import (
	"fmt"
	"io"
	"net/http"

	"github.com/bells307/everydaypic/internal/domain/entity"
	"github.com/bells307/everydaypic/internal/domain/usecase"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type imageHandler struct {
	imageUsecase usecase.ImageUsecase
}

func NewImageHandler(imageUsecase usecase.ImageUsecase) *imageHandler {
	return &imageHandler{imageUsecase: imageUsecase}
}

func (h *imageHandler) Register(e *gin.Engine) {
	images := e.Group("/image")
	{
		// images.GET("/")
		images.POST("/:name", h.createImage)
		images.DELETE("/:id", h.deleteImage)
	}
}

func (h *imageHandler) createImage(c *gin.Context) {
	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}

	uuid, err := uuid.NewRandom()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}

	img := entity.Image{
		ID:       uuid.String(),
		Name:     c.Param("name"),
		FileName: fmt.Sprintf("%s.jpg", uuid.String()),
	}

	err = h.imageUsecase.UploadImage(c.Request.Context(), img, data)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}

	c.JSON(200, uuid)
}

func (h *imageHandler) deleteImage(c *gin.Context) {
	id := c.Param("id")
	err := h.imageUsecase.DeleteImage(c.Request.Context(), id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
}
