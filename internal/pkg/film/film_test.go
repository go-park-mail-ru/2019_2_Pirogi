package film

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateNewFilm(t *testing.T) {
	newFilmInfo := models.FilmInfo{
		Title:       "title",
		Description: "description",
		Date:        "1998",
		Actors:      nil,
		Genres:      nil,
		Directors:   nil,
		Rating:      0,
		Image:       "",
		ReviewsNum:  models.ReviewsNum{},
	}

	newFilm := models.NewFilm{FilmInfo: newFilmInfo}
	expectedFilm := models.Film{
		ID:       0,
		FilmInfo: newFilmInfo,
	}
	newNewFilm, err := CreateNewFilm(0, &newFilm)
	require.Nil(t, err)
	require.Equal(t, expectedFilm, newNewFilm)

}
