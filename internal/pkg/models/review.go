package models

import "time"

type ReviewsNum int

type NewReview struct {
	Title    string `json:"title"`
	Body     string `json:"body"`
	FilmID   ID     `json:"film_id"`
	AuthorID ID     `json:"author_id, omitempty"`
}

// TODO: remove binary choice of film's like/dislike
type Review struct {
	NewReview
	ID ID `json:"id"`
	//TODO: заменить на гошный формат
	Date  time.Time `json:"date"`
	Likes int       `json:"likes"`
}
