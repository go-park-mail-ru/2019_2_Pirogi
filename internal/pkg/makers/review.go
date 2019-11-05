package makers

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"html"
	"time"
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
