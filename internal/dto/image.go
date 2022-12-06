package dto

import "io"

// Создание изображения
type CreateImage struct {
	Name     string
	FileName string
	UserID   string
	FileSize int64
	Data     io.ReadSeeker
}

// Получение изображений по фильтру
type GetImages struct {
	// ID изображения
	ID []string
	// Имя файла
	FileName []string
}
