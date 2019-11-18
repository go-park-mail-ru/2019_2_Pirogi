package files

import (
	"github.com/labstack/echo"
	"net/http"
	"os"
	"path/filepath"
)

func WriteFile(filename string, fileBytes []byte) error {
	newFile, err := os.Create(filepath.Join(filename))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "can not create file")
	}
	if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "can not open file for write")
	}
	return nil
}
