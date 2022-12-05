package dto

import "io"

// Загрузить файл
type UploadFile struct {
	Bucket   string
	Name     string
	Filename string
	FileSize int64
	Data     io.ReadSeeker
}
