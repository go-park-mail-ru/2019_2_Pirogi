package domains

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNew(t *testing.T) {
	expectedError := &Error{
		Status: 404,
		Error:  "Страница не найдена; Что делать?",
	}
	newError := NewError(404, "Страница не найдена", "Что делать?")
	require.Equal(t, expectedError, newError)
}

