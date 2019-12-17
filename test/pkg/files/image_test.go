package files

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/files"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDetectContentTypeOK(t *testing.T) {
	const testFileName = "./test.png"
	file, err := os.Open(testFileName)
	defer func() { _ = file.Close() }()
	require.NoError(t, err)
	image, err := ioutil.ReadAll(file)
	require.NoError(t, err)
	expectedEnding := ".png"
	actualEnding, e := files.DetectContentType(image)
	require.Nil(t, e)
	require.Equal(t, expectedEnding, actualEnding)
}

func TestDetectContentTypeFail(t *testing.T) {
	var image []byte
	_, e := files.DetectContentType(image)
	require.NotNil(t, e)
}
