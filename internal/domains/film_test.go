package domains

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/validation"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMakeTruncFilm(t *testing.T) {
	expected := testFilmTrunc
	actual := testFilm.Trunc()
	require.Equal(t, expected, actual)
}

func TestMakeFilm(t *testing.T) {
	expected := Film{
		ID:          2,
		Title:       testFilmNew.Title,
		Year:        testFilmNew.Year,
		Genres:      testFilmNew.Genres,
		Mark:        Mark(0),
		Description: testFilmNew.Description,
		Countries:   testFilmNew.Countries,
		PersonsID:   testFilmNew.PersonsID,
		Images:      []Image{"default.png"},
		ReviewsNum:  0,
	}
	actual := testFilmNew.ToFilm(2)
	require.Equal(t, expected, actual)
}

func TestMakeFullFilm(t *testing.T) {
	expected := testFilmFull
	actual := testFilm.Full([]Person{testPerson})
	require.Equal(t, expected, actual)
}

func TestPrepareModelNewFilm(t *testing.T) {
	validation.InitValidator()
	body, err := testFilmNew.MarshalJSON()
	require.NoError(t, err)
	model := FilmNew{}
	err = model.Make(body)
	require.NoError(t, err)
	require.Equal(t, testFilmNew, model)
}
