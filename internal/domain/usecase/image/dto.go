package image_usecase

type CreateImageDTO struct {
	Name     string
	Filename string
	Data     []byte
}
