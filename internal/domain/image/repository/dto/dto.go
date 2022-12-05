package dto

// Фильтр изображений
type GetImagesFilter struct {
	// ID изображения
	ID []string `form:"id" json:"id"`
	// Имя файла
	FileName []string `form:"fileName" json:"fileName"`
}
