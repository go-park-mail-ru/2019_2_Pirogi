package files

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadLinesInvalidFile(t *testing.T) {
	var invalidImage []byte
	err := WriteFile("", invalidImage)
	require.Error(t, err)
}

func TestReadLinesInvalidPath(t *testing.T) {
	const testFileName = "./test.png"
	const invalidPath = "./@@@@"
	file, err := os.Open(testFileName)
	defer func() { _ = file.Close() }()
	require.NoError(t, err)
	validImage, err := ioutil.ReadAll(file)
	require.NoError(t, err)
	err = WriteFile(invalidPath, validImage)
	require.Error(t, err)
}

func TestReadLines(t *testing.T) {
	const testFileName = "./test.png"
	const validPath = "./"
	const filename = "tmp.png"
	file, err := os.Open(testFileName)
	defer func() { _ = file.Close() }()
	require.NoError(t, err)
	validImage, err := ioutil.ReadAll(file)
	require.NoError(t, err)
	err = WriteFile(validPath+filename, validImage)
	require.NoError(t, err)
	err = os.Remove(validPath + filename)
	require.NoError(t, err)
}
