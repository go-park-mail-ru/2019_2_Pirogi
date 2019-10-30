package models

type Person struct {
	ID         ID      `json:"id"`
	Type       int     `json:"type"`
	Name       string  `json:"name"`
	Birthday   string  `json:"birthday"`
	Birthplace string  `json:"birthplace"`
	GenresID   []int   `json:"genres_id"`
	FilmsID    []int   `json:"films_id"`
	Rating     float32 `json:"rating"`
	ImagesID   []int   `json:"images_id"`
	AwardsID   []int   `json:"awards_id"`
}
