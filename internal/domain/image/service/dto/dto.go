package dto

import repoDTO "github.com/bells307/everydaypic/internal/domain/image/repository/dto"

type CreateImage struct {
	Name     string
	FileName string
	UserID   string
}

// Фильтр изображений
type GetImagesFilter struct {
	// ID изображения
	ID []string `form:"id" json:"id"`
	// Имя файла
	FileName []string `form:"fileName" json:"fileName"`
}

func (gif GetImagesFilter) ToRepoDTO() repoDTO.GetImagesFilter {
	return repoDTO.GetImagesFilter{
		ID:       gif.ID,
		FileName: gif.FileName,
	}
}
