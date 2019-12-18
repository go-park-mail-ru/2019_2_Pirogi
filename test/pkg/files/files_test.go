package files

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/files"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadLinesInvalidFile(t *testing.T) {
	var invalidImage []byte
	err := files.WriteFile("", invalidImage)
	require.NotNil(t, err)
}

func TestReadLinesInvalidPath(t *testing.T) {
	const testFileName = "./test.png"
	const invalidPath = "/etc/lib/~123'asdqw@@@@"
	file, err := os.Open(testFileName)
	defer func() { _ = file.Close() }()
	require.NoError(t, err)
	validImage, err := ioutil.ReadAll(file)
	require.NoError(t, err)
	e := files.WriteFile(invalidPath, validImage)
	require.NotNil(t, e)
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
	e := files.WriteFile(validPath+filename, validImage)
	require.Nil(t, e)
	err = os.Remove(validPath + filename)
	require.NoError(t, err)
}
