package v1

import (
	"net/http"

	"github.com/bells307/everydaypic/internal/domain/usecase"
	"github.com/gin-gonic/gin"
)

type imageHandler struct {
	imageUsecase usecase.ImageUsecase
}

func NewImageHandler(imageUsecase usecase.ImageUsecase) *imageHandler {
	return &imageHandler{imageUsecase: imageUsecase}
}

func (h *imageHandler) Register(e *gin.Engine) {
	images := e.Group("/images")
	{
		// images.GET("/")
		images.POST("/", h.createImage)
		images.DELETE("/:id", h.deleteImage)
	}
}

func (h *imageHandler) createImage(c *gin.Context) {
	var data []byte
	_, err := h.imageUsecase.UploadImage(c.Request.Context(), "img1", data)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
}

func (h *imageHandler) deleteImage(c *gin.Context) {
	id := c.Param("id")
	err := h.imageUsecase.DeleteImage(c.Request.Context(), id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
}
