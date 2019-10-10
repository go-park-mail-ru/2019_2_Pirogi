package common

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/images"
	"github.com/labstack/echo"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func WriteFile(fileBytes []byte, base string) (filename string, err error) {
	ending, e := images.DetectContentType(fileBytes)
	if e != nil {
		return "", echo.NewHTTPError(e.Status, e.Error)
	}
	filename = images.GenerateFilename(time.Now().String(), strconv.Itoa(rand.Int()), ending)
	newPath := filepath.Join(base, filename)
	newFile, err := os.Create(newPath)
	if err != nil {
		return "", echo.NewHTTPError(500, "can not create file")
	}
	if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
		return "", echo.NewHTTPError(500, "can not open file for write")
	}
	return filename, nil
}
