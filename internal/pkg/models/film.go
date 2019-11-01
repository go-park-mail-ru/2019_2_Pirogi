package models

// idk how to remove data duplication in case of new object
type NewFilm struct {
	Title         string        `json:"title" valid:"alphanum, stringlength(1|50)"`
	Description   string        `json:"description" valid:"alphanum, stringlength(8|50)"`
	Year          string        `json:"year" valid:"year"`
	Countries     []string      `json:"countries" valid:"countries"`
	Genres        []Genre       `json:"genres" valid:"genres"`
	Actors        []PersonTrunc `json:"actors" valid:"persons_trunc, optional"`
	Directors     []PersonTrunc `json:"directors" valid:"persons_trunc, optional"`
	Producers     []PersonTrunc `json:"producers" valid:"persons_trunc, optional"`
	Compositors   []PersonTrunc `json:"compositors" valid:"persons_trunc, optional"`
	Screenwriters []PersonTrunc `json:"screenwriters" valid:"persons_trunc, optional"`
	Poster        Image         `json:"poster" valid:"optional"`
	Images        []Image       `json:"image" valid:"images, optional"`
}

type Film struct {
	FilmTrunc   `valid:"required"`
	Description string        `json:"description" valid:"alphanum, stringlength(8|50)"`
	Countries   []string      `json:"countries" valid:"countries"`
	Actors      []PersonTrunc `json:"actors" valid:"persons_trunc, optional"`
	Directors   []PersonTrunc `json:"directors" valid:"persons_trunc, optional"`
	Images      []Image       `json:"image" valid:"optional"`
	ReviewsNum  int           `json:"reviews_num" valid:"numeric, optional"`
}

type FilmTrunc struct {
	ID     ID      `json:"id" bson:"_id" valid:"numeric,optional"`
	Title  string  `json:"title" valid:"alphanum, stringlength(1|50)"`
	Year   string  `json:"year" valid:"year, stringlength(4|4)"`
	Genres []Genre `json:"genres" valid:"genres"`
	Poster Image   `json:"poster" valid:"optional"`
	Mark   Mark    `json:"mark" valid:"mark"`
}
