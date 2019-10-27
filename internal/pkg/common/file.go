package common

import (
	"bufio"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

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
	newPath := filepath.Join(base, generatedFilename)
	newFile, err := os.Create(newPath)
	if err != nil {
		return "", echo.NewHTTPError(http.StatusInternalServerError, "can not create file")
	}
	if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
		return "", echo.NewHTTPError(http.StatusInternalServerError, "can not open file for write")
	}
	return generatedFilename, nil
}

func WriteBytes(data []byte, path string) error {
	newFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer newFile.Close()
	_, err = newFile.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func ReadLines(filename string) ([]string, error) {
	result := make([]string, 0)
	defer func() {
		if e := recover(); e != nil {
			println("recovered from ReadLines", e)
		}
	}()
	file, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result, nil
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
