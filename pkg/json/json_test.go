package json

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMakeJSONArray(t *testing.T) {
	input := [][]byte{{'a', 'b', 'c'}, {'d', 'e', 'f'}}
	expected := []byte{'[', 'a', 'b', 'c', ',', 'd', 'e', 'f', ']'}
	actual := MakeJSONArray(input)
	require.Equal(t, expected, actual)
}
