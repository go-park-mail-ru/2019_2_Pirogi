package models

// idk how to remove data duplication in case of new object
type NewFilm struct {
	Title         string        `json:"title"`
	Description   string        `json:"description"`
	Year          int           `json:"year"`
	Countries     []string      `json:"countries"`
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
	ReviewsNum  int           `json:"reviews_num"`
}

type FilmTrunc struct {
	ID     ID       `json:"id" bson:"_id"`
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Genres []Genre  `json:"genres"`
	Poster Image    `json:"poster"`
	Rating Mark `json:"rating"`
}
