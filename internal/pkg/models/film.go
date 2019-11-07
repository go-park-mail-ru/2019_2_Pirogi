package models

// idk how to remove data duplication in case of new object
type NewFilm struct {
	Title       string   `json:"title" valid:"text, stringlength(1|50)"`
	Description string   `json:"description" valid:"text, stringlength(8|50)"`
	Year        string   `json:"year" valid:"year"`
	Countries   []string `json:"countries" valid:"countries"`
	Genres      []Genre  `json:"genres" valid:"genres"`
	PersonsID   []ID     `json:"persons_id" valid:"ids, optional"`
	Trailer     string   `json:"trailer" valid:"link, optional"`
}

type Film struct {
	ID          ID       `json:"id" bson:"_id" valid:"numeric,optional"`
	Title       string   `json:"title" valid:"text, stringlength(1|50)"`
	Year        string   `json:"year" valid:"year, stringlength(4|4)"`
	Genres      []Genre  `json:"genres" valid:"genres"`
	Mark        Mark     `json:"mark" valid:"mark"`
	Description string   `json:"description" valid:"text, stringlength(8|50)"`
	Countries   []string `json:"countries" valid:"countries"`
	PersonsID   []ID     `json:"persons_id" valid:"ids, optional"`
	Images      []Image  `json:"images" valid:"images, optional"`
	ReviewsNum  int      `json:"reviews_num" valid:"numeric, optional"`
	Trailer     string   `json:"trailer" valid:"link, optional"`
}

type FilmTrunc struct {
	ID     ID      `json:"id" valid:"numeric,optional"`
	Title  string  `json:"title" valid:"text, stringlength(1|50)"`
	Year   string  `json:"year" valid:"year, stringlength(4|4)"`
	Genres []Genre `json:"genres" valid:"genres"`
	Mark   Mark    `json:"mark" valid:"mark"`
	Poster Image   `json:"poster" valid:"image"`
}

type FilmFull struct {
	ID          ID            `json:"id" bson:"_id" valid:"numeric,optional"`
	Title       string        `json:"title" valid:"text, stringlength(1|50)"`
	Year        string        `json:"year" valid:"year, stringlength(4|4)"`
	Genres      []Genre       `json:"genres" valid:"genres"`
	Mark        Mark          `json:"mark" valid:"mark"`
	Description string        `json:"description" valid:"text, stringlength(8|50)"`
	Countries   []string      `json:"countries" valid:"countries"`
	Persons     []PersonTrunc `json:"persons" valid:"optional"`
	Images      []Image       `json:"images" valid:"images, optional"`
	ReviewsNum  int           `json:"reviews_num" valid:"numeric, optional"`
	Trailer     string        `json:"trailer" valid:"link, optional"`
}
