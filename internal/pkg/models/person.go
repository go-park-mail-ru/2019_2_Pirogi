package models

type Person struct {
	ID         ID          `json:"id"`
	Roles      []Role      `json:"type"`
	Name       string      `json:"name"`
	Birthday   string      `json:"birthday"`
	Birthplace string      `json:"birthplace"`
	Genres     []Genre     `json:"genres_id"`
	Films      []FilmTrunc `json:"films_id"`
	Likes      int         `json:"rating"`
	Images     []Image     `json:"images_id"`
	Awards     []Award     `json:"awards_id"`
}

type PersonTrunc struct {
	ID   ID     `json:"id"`
	Name string `json:"name"`
}
