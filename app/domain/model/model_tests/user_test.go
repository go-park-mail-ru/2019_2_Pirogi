package model_tests

import (
	model2 "github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/validation"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPrepareModelNewUser(t *testing.T) {
	validation.InitValidator()
	userNew := model2.UserNew{}
	userNew.Create("i@artbakulev.com", "1234567890", "Artyom")
	body, err := userNew.MarshalJSON()
	require.NoError(t, err)
	model := model2.UserNew{}
	err = model.Make(body)
	require.NoError(t, err)
	require.Equal(t, userNew, model)
}
