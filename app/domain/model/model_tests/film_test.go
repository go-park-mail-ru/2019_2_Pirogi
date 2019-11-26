package model_tests

import (
	"testing"

	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/validation"
	"github.com/stretchr/testify/require"
)

func TestMakeTruncFilm(t *testing.T) {
	expected := testFilmTrunc
	actual := testFilm.Trunc()
	require.Equal(t, expected, actual)
}

func TestMakeFilm(t *testing.T) {
	expected := model.Film{
		ID:          2,
		Title:       testFilmNew.Title,
		Year:        testFilmNew.Year,
		Genres:      testFilmNew.Genres,
		Mark:        model.Mark(0),
		Description: testFilmNew.Description,
		Countries:   testFilmNew.Countries,
		PersonsID:   testFilmNew.PersonsID,
		Images:      []model.Image{"default.png"},
		ReviewsNum:  0,
	}
	actual := testFilmNew.ToFilm(2)
	require.Equal(t, expected, actual)
}

func TestMakeFullFilm(t *testing.T) {
	expected := testFilmFull
	actual := testFilm.Full([]model.Person{testPerson})
	require.Equal(t, expected, actual)
}

func TestPrepareModelNewFilm(t *testing.T) {
	validation.InitValidator()
	body, err := testFilmNew.MarshalJSON()
	require.NoError(t, err)
	filmNew := model.FilmNew{}
	err = filmNew.Make(body)
	require.NoError(t, err)
	require.Equal(t, testFilmNew, filmNew)
}
