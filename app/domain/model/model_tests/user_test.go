package model_tests

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/test/fixture"
	"testing"

	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/validation"
	"github.com/stretchr/testify/require"
)

func TestPrepareModelNewUser(t *testing.T) {
	validation.InitValidator()
	userNew := fixture.UserNew
	body, err := userNew.MarshalJSON()
	require.NoError(t, err)
	model := model.UserNew{}
	err = model.Make(body)
	require.NoError(t, err)
	require.Equal(t, userNew, model)
}
