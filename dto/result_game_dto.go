package dto

type ResultGameDTO struct {
	Results []struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Genres []struct {
			Name string `json:"name"`
		} `json:"genres"`
	} `json:"results"`
}
