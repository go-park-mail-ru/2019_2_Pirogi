package models

// idk how to remove data duplication in case of new object
type NewFilm struct {
	Title         string        `json:"title"`
	Description   string        `json:"description"`
	Date          string        `json:"date"`
	Country       string        `json:"country"`
	Genres        []Genre       `json:"genres"`
	Actors        []PersonTrunc `json:"actors"`
	Directors     []PersonTrunc `json:"directors"`
	Producers     []PersonTrunc `json:"producers"`
	Compositors   []PersonTrunc `json:"compositors"`
	Screenwriters []PersonTrunc `json:"screenwriters"`
	Poster        Image         `json:"poster"`
	Images        []Image       `json:"image"`
}

type Film struct {
	FilmTrunc
	Description string        `json:"description"`
	Actors      []PersonTrunc `json:"actors"`
	Directors   []PersonTrunc `json:"directors"`
	Images      []Image       `json:"image"`
	ReviewsNum
}

type FilmTrunc struct {
	ID     ID       `json:"id" bson:"_id"`
	Title  string   `json:"title"`
	Date   string   `json:"date"`
	Genres []Genre  `json:"genres"`
	Poster Image    `json:"poster"`
	Rating FilmMark `json:"rating"`
}
