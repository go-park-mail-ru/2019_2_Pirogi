package handlers

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/domains"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/infrastructure/database"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/common"
	"github.com/labstack/echo"
)

func GetImagesHandler(conn database.Database) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		user, ok := auth.GetUserByRequest(ctx.Request(), conn)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "no auth")
		}

		fileBytes, err := ParseRequestAndWriteFile(ctx)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		var base string
		switch {
		case strings.Contains(ctx.Request().URL.Path, "users"):
			base = configs.Default.UsersImageUploadPath
		case strings.Contains(ctx.Request().URL.Path, "films"):
			base = configs.Default.FilmsImageUploadPath
		case strings.Contains(ctx.Request().URL.Path, "persons"):
			base = configs.Default.PersonsImageUploadPath
		default:
			return echo.NewHTTPError(http.StatusBadRequest, "wrong path")
		}
		filename, err := common.WriteFileWithGeneratedName(fileBytes, base)
		if err != nil {
			return err
		}

		// TODO: разобраться с изображениями
		user.Image = domains.Image(filename)
		e := conn.Upsert(user)
		if e != nil {
			return echo.NewHTTPError(e.Status, e.Error)
		}
		return nil
	}
}

func ParseRequestAndWriteFile(ctx echo.Context) ([]byte, error) {
	ctx.Request().Body = http.MaxBytesReader(ctx.Response(), ctx.Request().Body, files.MaxUploadSize)
	if err := ctx.Request().ParseMultipartForm(files.MaxUploadSize); err != nil {
		return []byte{}, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	defer ctx.Request().Body.Close()

	file, _, err := ctx.Request().FormFile("file")
	if err != nil {
		return []byte{}, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	defer func() {
		err := file.Close()
		if err != nil {
			ctx.Logger().Warn("can not close file")
		}
	}()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return []byte{}, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return fileBytes, nil
}