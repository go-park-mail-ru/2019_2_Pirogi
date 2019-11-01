package models

type NewPerson struct {
	Name       string      `json:"name"`
	Roles      []Role      `json:"type"`
	Birthday   string      `json:"birthday"`
	Birthplace string      `json:"birthplace"`
	Genres     []Genre     `json:"genres_id"`
	Films      []FilmTrunc `json:"films_id"`
	Likes      int         `json:"rating, omitempty"`
	Images     []Image     `json:"images_id"`
}

type Person struct {
	PersonTrunc
	Roles      []Role      `json:"type"`
	Birthday   string      `json:"birthday"`
	Birthplace string      `json:"birthplace"`
	Genres     []Genre     `json:"genres_id"`
	Films      []FilmTrunc `json:"films_id"`
	Likes      int         `json:"rating, omitempty"`
	Images     []Image     `json:"images_id"`
}

type PersonTrunc struct {
	ID   ID     `json:"id, omitempty"`
	Name string `json:"name"`
}
