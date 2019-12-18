package model_tests

import (
	"testing"

	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/validation"
	"github.com/stretchr/testify/require"
)

func TestMakeTruncPerson(t *testing.T) {
	expected := testPersonTrunc
	actual := testPerson.Trunc()
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
