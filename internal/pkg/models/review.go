package models

import "time"

type NewReview struct {
	Title  string `json:"title" valid:"title, stringlength(2|50)"`
	Body   string `json:"body" valid:"description, stringlength(8|50)"`
	FilmID ID     `json:"film_id" valid:"numeric"`
	//TODO: убрать отсюда автор ID
	AuthorID ID `json:"author_id, omitempty" valid:"numeric, optional"`
}

// TODO: remove binary choice of film's like/dislike
type Review struct {
	ID       ID        `json:"id, omitempty" bson:"_id" valid:"required"`
	Title    string    `json:"title" valid:"title, stringlength(2|50)"`
	Body     string    `json:"body" valid:"description, stringlength(8|50)"`
	FilmID   ID        `json:"film_id" valid:"numeric"`
	AuthorID ID        `json:"author_id, omitempty" valid:"numeric"`
	Date     time.Time `json:"date" valid:"time"`
	Likes    int       `json:"likes" valid:"numeric, optional"`
}
