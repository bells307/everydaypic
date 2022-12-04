package dto

import "io"

type CreateImage struct {
	Name     string
	FileName string
	UserID   string
	FileSize int64
	Data     io.ReadSeeker
}

// Фильтр изображений
type GetImages struct {
	// ID изображения
	ID []string `form:"id" json:"id"`
	// Имя файла
	FileName []string `form:"fileName" json:"fileName"`
}
