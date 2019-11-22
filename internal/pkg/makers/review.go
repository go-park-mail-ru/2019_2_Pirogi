package makers

import (
	"html"
	"time"

	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
)

func MakeReview(id models.ID, in models.NewReview) models.Review {
	return models.Review{
		ID:       id,
		Date:     time.Now(),
		Likes:    0,
		Title:    html.EscapeString(in.Title),
		Body:     html.EscapeString(in.Body),
		FilmID:   in.FilmID,
		AuthorID: in.AuthorID,
	}
}

func MakeReviewFull(in models.Review, author models.UserTrunc, mark models.Mark) models.ReviewFull {
	return models.ReviewFull{
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
