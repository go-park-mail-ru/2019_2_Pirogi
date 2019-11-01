package models

import "time"

type NewReview struct {
	Title    string `json:"title" valid:"title, stringlength(8|50)"`
	Body     string `json:"body" valid:"description, stringlength(8|50)"`
	FilmID   ID     `json:"film_id" valid:"numeric"`
	AuthorID ID     `json:"author_id, omitempty" valid:"numeric"`
}

// TODO: remove binary choice of film's like/dislike
type Review struct {
	NewReview `valid:"required"`
	ID        ID        `json:"id, omitempty" bson:"_id"`
	Date      time.Time `json:"date" valid:"time"`
	Likes     int       `json:"likes" valid:"numeric, optional"`
}
