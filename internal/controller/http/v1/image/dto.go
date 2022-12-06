package image

import "mime/multipart"

// Создание изображения
type CreateImage struct {
	// Имя изображения
	Name string `form:"name" binding:"required"`
	// Имя файла
	FileName string `form:"fileName" binding:"required"`
	// Бинарные данные файла
	File *multipart.FileHeader `form:"file" binding:"required" swaggertype:"string" format:"binary"`
}

// Фильтр изображений
type GetImages struct {
	// ID изображения
	ID []string `form:"id" json:"id"`
	// Имя файла
	FileName []string `form:"fileName" json:"fileName"`
}
