package json

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMakeJSONArray(t *testing.T) {
	input := [][]byte{{'a', 'b', 'c'}, {'d', 'e', 'f'}}
	expected := []byte{'[', 'a', 'b', 'c', ',', 'd', 'e', 'f', ']'}
	actual := json.MakeJSONArray(input)
	require.Equal(t, expected, actual)
}
