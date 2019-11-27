package model_tests

import (
	"testing"

	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/configuration"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/validation"
	"github.com/stretchr/testify/require"
)

func TestMakeTruncPerson(t *testing.T) {
	expected := testPersonTrunc
	actual := testPerson.Trunc()
	require.Equal(t, expected, actual)
}

func TestMakePerson(t *testing.T) {
	err := configuration.UnmarshalConfigs("../../configs")
	require.NoError(t, err)
	expected := model.Person{
		ID:         testPerson.ID,
		Name:       testPerson.Name,
		Mark:       0,
		Roles:      testPerson.Roles,
		Birthday:   testPerson.Birthday,
		Birthplace: testPerson.Birthplace,
		Genres:     []model.Genre{},
		FilmsID:    []model.ID{},
		Likes:      0,
		Images:     []model.Image{"default.png"},
	}
	actual := testPersonNew.ToPerson(2)
	require.Equal(t, expected, actual)
}

func TestMakeFullPerson(t *testing.T) {
	expected := testPersonFull
	actual := testPerson.Full([]model.Film{testFilm})
	require.Equal(t, expected, actual)
}

func TestPrepareModelNewPerson(t *testing.T) {
	validation.InitValidator()
	body, err := testPersonNew.MarshalJSON()
	require.NoError(t, err)
	personNew := model.PersonNew{}
	err = personNew.Make(body)
	require.NoError(t, err)
	require.Equal(t, testPersonNew, personNew)
}
