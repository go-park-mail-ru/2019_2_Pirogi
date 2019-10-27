package common

import (
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"testing"
)

func TestReadLinesInvalidFile(t *testing.T) {
	var invalidImage []byte
	_, err := WriteFileWithGeneratedName(invalidImage, "")
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
	_, err = WriteFileWithGeneratedName(validImage, invalidPath)
	require.Error(t, err)
}

func TestReadLines(t *testing.T) {
	const testFileName = "./test.png"
	const validPath = "./"
	file, err := os.Open(testFileName)
	defer func() { _ = file.Close() }()
	require.NoError(t, err)
	validImage, err := ioutil.ReadAll(file)
	require.NoError(t, err)
	name, err := WriteFileWithGeneratedName(validImage, validPath)
	require.NoError(t, err)
	require.NotNil(t, name)
}
