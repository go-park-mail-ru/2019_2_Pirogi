package models

type ReviewsNum struct {
	Total    int `json:"total"`
	Positive int `json:"positive"`
	Negative int `json:"negative"`
}

type FilmInfo struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Date        string  `json:"date"`
	Rating      float32 `json:"rating"`
	ActorsID    []ID    `json:"actors"`
	GenresID    []ID    `json:"genres"`
	DirectorsID []ID    `json:"directors"`
	ImagesID    []ID    `json:"image"`
	ReviewsNum
}

type NewFilm struct {
	FilmInfo
}

type Film struct {
	ID ID `json:"id" bson:"_id"`
	FilmInfo
}
