package v1

import "mime/multipart"

// Создание изображения
type CreateImage struct {
	// Имя изображения
	Name string `form:"name" binding:"required"`
	// Имя файла
	Filename string `form:"filename" binding:"required"`
	// Бинарные данные файла
	File *multipart.FileHeader `form:"file" binding:"required"`
}
