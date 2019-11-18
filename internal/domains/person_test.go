package domains

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/domains/film"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/configuration"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/validation"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMakeTruncPerson(t *testing.T) {
	expected := testPersonTrunc
	actual := testPerson.Trunc()
	require.Equal(t, expected, actual)
}

func TestMakePerson(t *testing.T) {
	err := configuration.UnmarshalConfigs("../../configs")
	require.NoError(t, err)
	expected := Person{
		ID:         testPerson.ID,
		Name:       testPerson.Name,
		Mark:       0,
		Roles:      testPerson.Roles,
		Birthday:   testPerson.Birthday,
		Birthplace: testPerson.Birthplace,
		Genres:     []Genre{},
		FilmsID:    []ID{},
		Likes:      0,
		Images:     []Image{"default.png"},
	}
	actual := testPersonNew.ToPerson(2)
	require.Equal(t, expected, actual)
}

func TestMakeFullPerson(t *testing.T) {
	expected := testPersonFull
	actual := testPerson.Full([]film.Film{testFilm})
	require.Equal(t, expected, actual)
}

func TestPrepareModelNewPerson(t *testing.T) {
	validation.InitValidator()
	body, err := testPersonNew.MarshalJSON()
	require.NoError(t, err)
	model := PersonNew{}
	err = model.Make(body)
	require.NoError(t, err)
	require.Equal(t, testPersonNew, model)
}
