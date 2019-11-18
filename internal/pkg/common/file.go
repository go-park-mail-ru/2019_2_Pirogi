package common

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"gopkg.in/yaml.v2"

	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"

	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/images"
	"github.com/labstack/echo"
)

func WriteFileWithGeneratedName(fileBytes []byte, base string) (generatedFilename string, err error) {
	ending, e := images.DetectContentType(fileBytes)
	if e != nil {
		return "", echo.NewHTTPError(e.Status, e.Error)
	}
	generatedFilename = images.GenerateFilename(time.Now().String(), strconv.Itoa(rand.Int()), ending)
	newFile, err := os.Create(filepath.Join(base, generatedFilename))
	if err != nil {
		return "", echo.NewHTTPError(http.StatusInternalServerError, "can not create file")
	}
	if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
		return "", echo.NewHTTPError(http.StatusInternalServerError, "can not open file for write")
	}
	return generatedFilename, nil
}

func NormalizePath(path *string) {
	if (*path)[len(*path)-1] != '/' {
		*path += "/"
	}
}

func UnmarshalConfigs(configsPath *string) error {
	NormalizePath(configsPath)
	file, err := ioutil.ReadFile(*configsPath + "default.yaml")
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(file, &configs.Default)
	if err != nil {
		return err
	}

	file, err = ioutil.ReadFile(*configsPath + "headers.yaml")
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(file, &configs.Headers)
	if err != nil {
		return err
	}
	return nil
}
