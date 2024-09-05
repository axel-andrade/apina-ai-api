package models

type Content struct {
	Base        `bson:",inline"`
	Name        string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`
	Kind        string `bson:"kind" json:"kind"`
	File        []byte `bson:"file" json:"file"`
	Filename    string `bson:"filename" json:"filename"`
	Text        string `bson:"text" json:"text"`
	Status      string `bson:"status" json:"status"`
	Language    string `bson:"language" json:"language"`
}
