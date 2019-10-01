package film

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	user2 "github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/user"
)

func CreateNewFilm(title, description string, genres, actors, directors []string) (models.Film, error) {
	film := models.Film{
		Title:       title,
		Description: description,
		Image:       user2.GetMD5Hash(title + description),
		Genres:      genres,
		Actors:      actors,
		Directors:   directors,
		Rating:      0,
		ReviewsNum: models.ReviewsNum{
			Total:    0,
			Positive: 0,
			Negative: 0,
		},
	}
	return film, nil
}
