package v1

import "mime/multipart"

type CreateImage struct {
	Name     string                `form:"name" binding:"required"`
	Filename string                `form:"filename" binding:"required"`
	File     *multipart.FileHeader `form:"file" binding:"required"`
}
