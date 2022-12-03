package image

import "mime/multipart"

// Создание изображения
type CreateImage struct {
	// Имя изображения
	Name string `form:"name" binding:"required"`
	// Имя файла
	FileName string `form:"fileName" binding:"required"`
	// Бинарные данные файла
	File *multipart.FileHeader `form:"file" binding:"required"`
}
