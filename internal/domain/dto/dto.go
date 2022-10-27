package dto

type CreateImage struct {
	Name     string
	FileName string
	Data     []byte
}

type GetImages struct {
	ID       []string `form:"id" json:"id"`
	Name     []string `form:"name" json:"name"`
	FileName []string `form:"filename" json:"filename"`
}
