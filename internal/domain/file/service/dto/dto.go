package dto

import (
	"io"

	"github.com/bells307/everydaypic/internal/domain/file/repository/dto"
)

// Загрузить файл
type UploadFile struct {
	Bucket   string
	Name     string
	Filename string
	FileSize int64
	Data     io.ReadSeeker
}

func (uf UploadFile) ToRepoDTO() dto.UploadFile {
	return dto.UploadFile{
		Bucket:   uf.Bucket,
		Name:     uf.Name,
		Filename: uf.Filename,
		FileSize: uf.FileSize,
		Data:     uf.Data,
	}
}
