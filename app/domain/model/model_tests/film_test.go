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
