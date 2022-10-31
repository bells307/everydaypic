package entity

type Image struct {
	ID       string         `json:"id" bson:"_id"`
	Filename string         `json:"filename" bson:"filename"`
	Metadata map[string]any `json:"metadata" bson:"metadata"`
}
