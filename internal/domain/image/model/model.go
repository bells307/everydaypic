package model

import "time"

// Изображение
type Image struct {
	// ID изображения
	ID string `json:"id" bson:"_id"`
	// Имя изображения
	Name string `json:"name" bson:"name"`
	// Имя файла
	FileName string `json:"fileName" bson:"fileName"`
	// ID пользователя, добавившего изображение
	UserID string `json:"userID" bson:"userID"`
	// Дата создания
	Created time.Time `json:"created" bson:"created"`
}
