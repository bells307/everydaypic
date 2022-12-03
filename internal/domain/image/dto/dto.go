package dto

import "io"

type CreateImage struct {
	Name     string
	FileName string
	UserID   string
	FileSize int64
	Data     io.Reader
}

type GetImages struct {
	ID       []string `form:"id" json:"id"`
	FileName []string `form:"fileName" json:"fileName"`
}
