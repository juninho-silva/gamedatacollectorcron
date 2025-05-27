package model

type Game struct {
	ID     string `json:"id" bson:"id"`
	Nome   string `json:"name" bson:"name"`
	Genero string `json:"genre" bson:"genre"`
}
