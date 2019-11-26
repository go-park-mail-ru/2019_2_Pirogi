package model_tests

import (
	"testing"

	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	expectedError := &model.Error{
		Status: 404,
		Error:  "Страница не найдена; Что делать?",
	}
	newError := model.NewError(404, "Страница не найдена", "Что делать?")
	require.Equal(t, expectedError, newError)
}
