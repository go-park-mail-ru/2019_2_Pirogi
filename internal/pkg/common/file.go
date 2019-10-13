package common

import (
	"bufio"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"

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
		return "", echo.NewHTTPError(500, "can not create file")
	}
	if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
		return "", echo.NewHTTPError(500, "can not open file for write")
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
