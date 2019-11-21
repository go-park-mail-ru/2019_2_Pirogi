package model

import (
	"github.com/asaskevich/govalidator"
	"golang.org/x/net/html"
	"time"
)

type ReviewNew struct {
	Title  string `json:"title" valid:"title, stringlength(2|50)"`
	Body   string `json:"body" valid:"description, stringlength(8|50)"`
	FilmID ID     `json:"film_id" valid:"numeric"`
	//TODO: убрать отсюда автор ID
	AuthorID ID `json:"author_id, omitempty" valid:"numeric, optional"`
}

func (nr *ReviewNew) ToReview(id ID) Review {
	return Review{
		ID:       id,
		Date:     time.Now(),
		Likes:    0,
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
	Likes    int       `json:"likes" valid:"numeric, optional"`
	Mark     Mark      `json:"mark" valid:"numeric, optional"`
}

type ReviewFull struct {
	ID     ID        `json:"id, omitempty" bson:"_id" valid:"required"`
	Title  string    `json:"title" valid:"title, stringlength(2|50)"`
	Body   string    `json:"body" valid:"description"`
	FilmID ID        `json:"film_id" valid:"numeric"`
	Author UserTrunc `json:"author, omitempty" valid:"numeric"`
	Date   time.Time `json:"date" valid:"time"`
	Likes  int       `json:"likes" valid:"numeric, optional"`
	Mark   Mark      `json:"stars" valid:"numeric, optional"`
}

func (r *Review) AddLike() {
	r.Likes += 1
}

func (r *Review) RemoveLike() {
	r.Likes -= 1
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
		Likes:  r.Likes,
		Mark:   r.Mark,
	}
}
