package entity

import "time"

// Изображение
type Image struct {
	ID         string         `json:"id" bson:"_id"`
	Filename   string         `json:"filename" bson:"filename"`
	Metadata   map[string]any `json:"metadata" bson:"metadata"`
	UploadDate time.Time      `json:"uploadDate" bson:"uploadDate"`
}

// Изображение дня
type DayImage struct {
	// Изображение
	Image
	// Дата установки
	SetAt time.Time `json:"setAt" bson:"setAt"`
}
