package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	"golang.org/x/net/html"
)

type ReviewNew struct {
	Title  string `json:"title" valid:"title, stringlength(2|50)"`
	Body   string `json:"body" valid:"description, stringlength(8|50)"`
	FilmID ID     `json:"film_id, omitempty" valid:"numeric, optional"`
	//TODO: убрать отсюда автор ID
	AuthorID ID `json:"author_id, omitempty" valid:"numeric, optional"`
}

func (nr *ReviewNew) ToReview(id ID) Review {
	return Review{
		ID:       id,
		Date:     time.Now(),
		Title:    html.EscapeString(nr.Title),
		Body:     html.EscapeString(nr.Body),
		FilmID:   nr.FilmID,
		AuthorID: nr.AuthorID,
	}
}

func (nr *ReviewNew) Make(body []byte) error {
	err := nr.UnmarshalJSON(body)
	if err != nil {
		return err
	}
	_, err = govalidator.ValidateStruct(nr)
	return err
}

type Review struct {
	ID       ID        `json:"id, omitempty" bson:"_id" valid:"required"`
	Title    string    `json:"title" valid:"title, stringlength(2|50)"`
	Body     string    `json:"body" valid:"description"`
	FilmID   ID        `json:"film_id" valid:"numeric"`
	AuthorID ID        `json:"author_id, omitempty" valid:"numeric"`
	Date     time.Time `json:"date" valid:"time"`
	Mark     Mark      `json:"mark" valid:"float, optional"`
}

type ReviewFull struct {
	ID     ID        `json:"id, omitempty" bson:"_id" valid:"required"`
	Title  string    `json:"title" valid:"title, stringlength(2|50)"`
	Body   string    `json:"body" valid:"description"`
	FilmID ID        `json:"film_id" valid:"numeric"`
	Author UserTrunc `json:"author, omitempty" valid:"numeric"`
	Date   time.Time `json:"date" valid:"time"`
	Mark   Mark      `json:"stars" valid:"numeric, optional"`
}

func (r *Review) SetMark(mark Mark) {
	r.Mark = mark
}

func (r *Review) Full(author UserTrunc) ReviewFull {
	return ReviewFull{
		ID:     r.ID,
		Title:  r.Title,
		Body:   r.Body,
		FilmID: r.FilmID,
		Author: author,
		Date:   r.Date,
		Mark:   r.Mark,
	}
}
