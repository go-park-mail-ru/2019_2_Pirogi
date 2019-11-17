package error

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/domains"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	expectedError := &domains.Error{
		Status: 404,
		Error:  "Страница не найдена; Что делать?",
	}
	newError := New(404, "Страница не найдена", "Что делать?")
	require.Equal(t, expectedError, newError)
}
