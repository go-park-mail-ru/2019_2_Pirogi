package makers

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/domains"
	"html"
	"time"
)

func MakeReview(id domains.ID, in domains.NewReview) domains.Review {
	return domains.Review{
		ID:       id,
		Date:     time.Now(),
		Likes:    0,
		Title:    html.EscapeString(in.Title),
		Body:     html.EscapeString(in.Body),
		FilmID:   in.FilmID,
		AuthorID: in.AuthorID,
	}
}

func MakeReviewFull(in domains.Review, author domains.UserTrunc, mark domains.Mark) domains.ReviewFull {
	return domains.ReviewFull{
		ID:     in.ID,
		Title:  in.Title,
		Body:   in.Body,
		FilmID: in.FilmID,
		Author: author,
		Date:   in.Date,
		Likes:  in.Likes,
		Mark:   mark,
	}
}
