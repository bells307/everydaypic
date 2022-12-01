package model

import "time"

// Изображение
type Image struct {
	ID       string    `json:"id" bson:"_id"`
	FileName string    `json:"filename" bson:"fileName"`
	UserID   string    `json:"userID" bson:"userID"`
	Created  time.Time `json:"created" bson:"created"`
}
