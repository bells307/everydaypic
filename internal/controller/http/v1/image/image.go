package image

import (
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/bells307/everydaypic/docs"
	"github.com/bells307/everydaypic/internal/dto"
	"github.com/bells307/everydaypic/internal/usecase"
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
			images.GET("/info", h.getImagesInfo)
			images.POST("", h.createImage)
			images.GET("/:id", h.downloadImage)
			// images.DELETE("/:id", h.deleteImage)
		}
	}
}

// @Summary     Получить информация об изображениях
// @Tags        image
// @ID          get-images-info
// @Accept      json
// @Produce     json
// @Param       input query dto.GetImages true "Фильтр изображений"
// @Success     200 {array} model.Image "Информация об изображении"
// @Failure     400 {string} string "Неправильно сформирован запрос"
// @Failure     404 {string} string "Изображение не найдено"
// @Failure     500 {string} string "Внутренняя ошибка сервиса"
// @Router      /v1/image/info [get]
func (h *imageHandler) getImagesInfo(c *gin.Context) {
	var getImages GetImages
	if err := c.Bind(&getImages); err != nil {
		return
	}

	dto := dto.GetImages{
		ID:       getImages.ID,
		FileName: getImages.FileName,
	}

	imgs, err := h.imageUsecase.GetImages(c.Request.Context(), dto)
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

// @Summary     Добавить изображение
// @Tags        image
// @ID          create-image
// @Accept      mpfd
// @Produce     json
// @Param       input formData CreateImage true "данные изображения"
// @Success     200 {array} model.Image "Изображение успешно создано"
// @Failure     400 {string} string "Неправильно сформирован запрос"
// @Failure     500 {string} string "Внутренняя ошибка сервиса"
// @Router      /v1/image [post]
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
		// TODO: поменять, когда будет реализована авторизация
		UserID:   "00000000-0000-0000-0000-000000000000",
		FileSize: file.Size,
		Data:     data,
	}

	img, err := h.imageUsecase.CreateImage(c.Request.Context(), dto)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, img)
}

// @Summary     Скачать изображение
// @Tags        image
// @ID          download-image
// @Produce     png,jpeg,gif,octet-stream
// @Param       id path string true "ID изображения"
// @Success     200 {string} binary "Бинарные данные изображения"
// @Failure     404 {string} string "Изображение не найдено"
// @Failure     500 {string} string "Внутренняя ошибка сервиса"
// @Router      /v1/image/{id} [get]
func (h *imageHandler) downloadImage(c *gin.Context) {
	id := c.Param("id")

	exists, err := h.imageUsecase.CheckImageExists(c.Request.Context(), id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if !exists {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	url, err := h.imageUsecase.GetDownloadUrl(c.Request.Context(), id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Redirect(http.StatusFound, url.String())
}
