package files

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func WriteFile(filename string, fileBytes []byte) *model.Error {
	newFile, err := os.Create(filepath.Join(filename))
	if err != nil {
		return model.NewError(http.StatusInternalServerError, "can not create file")
	}
	if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
		return model.NewError(http.StatusInternalServerError, "can not open file for write")
	}
	return nil
}

func ReadFileByteFromContext(ctx echo.Context) ([]byte, *model.Error) {
	ctx.Request().Body = http.MaxBytesReader(ctx.Response(), ctx.Request().Body, configs.Default.MaxFileUploadSize)
	if err := ctx.Request().ParseMultipartForm(configs.Default.MaxFileUploadSize); err != nil {
		return []byte{}, model.NewError(400, err.Error())
	}
	defer ctx.Request().Body.Close()

	file, _, err := ctx.Request().FormFile("file")
	if err != nil {
		return []byte{}, model.NewError(400, err.Error())
	}
	defer func() {
		err := file.Close()
		if err != nil {
			zap.S().Warn("can not close file")
		}
	}()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return []byte{}, model.NewError(400, err.Error())
	}
	return fileBytes, nil
}