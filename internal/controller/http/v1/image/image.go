package image

import (
	"net/http"

	"github.com/bells307/everydaypic/internal/domain/image/dto"
	"github.com/bells307/everydaypic/internal/domain/image/usecase"
	"github.com/gin-gonic/gin"
)

type imageHandler struct {
	imageUsecase *usecase.ImageUsecase
}

func NewImageHandler(imageUsecase *usecase.ImageUsecase) *imageHandler {
	return &imageHandler{imageUsecase}
}

func (h *imageHandler) Register(e *gin.Engine) {
	v1 := e.Group("/v1")
	{
		images := v1.Group("/image")
		{
			images.GET("", h.getImages)
			images.POST("", h.createImage)
			// images.GET("/:id/download", h.downloadImage)
			// images.DELETE("/:id", h.deleteImage)
		}
	}
}

// Получить изображения
func (h *imageHandler) getImages(c *gin.Context) {
	var getImages dto.GetImages
	if err := c.Bind(&getImages); err != nil {
		return
	}

	imgs, err := h.imageUsecase.GetImages(c.Request.Context(), getImages)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if len(imgs) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, imgs)
}

// Добавить изображение
func (h *imageHandler) createImage(c *gin.Context) {
	var createImage CreateImage
	if err := c.Bind(&createImage); err != nil {
		return
	}

	file := createImage.File
	data, err := file.Open()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer data.Close()

	dto := dto.CreateImage{
		Name:     createImage.Name,
		FileName: createImage.FileName,
		UserID:   "00000000-0000-0000-0000-000000000000",
		FileSize: file.Size,
		Data:     data,
	}

	img, err := h.imageUsecase.AddImage(c.Request.Context(), dto)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, img)
}
