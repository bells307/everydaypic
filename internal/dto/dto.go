package dto

// Создать изображение
type CreateImage struct {
	Name     string
	FileName string
	Data     []byte
}

// Получить изображение
type GetImages struct {
	ID       []string `form:"id" json:"id"`
	FileName []string `form:"filename" json:"filename"`
}
