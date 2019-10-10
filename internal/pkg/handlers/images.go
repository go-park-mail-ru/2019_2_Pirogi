package handlers

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/auth"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/common"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/database"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/images"
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetImagesHandler(conn database.Database) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		user, ok := auth.GetUserByRequest(ctx.Request(), conn)
		if !ok {
			return echo.NewHTTPError(401, "no auth")
		}

		fileBytes, err := ParseRequestAndWriteFile(ctx)
		if err != nil {
			return echo.NewHTTPError(500, err.Error())
		}

		var base string
		switch {
		case strings.Contains(ctx.Request().URL.Path, "users"):
			base = configs.UsersImageUploadPath
		case strings.Contains(ctx.Request().URL.Path, "films"):
			base = configs.FilmsImageUploadPath
		default:
			return echo.NewHTTPError(400, "wrong path")
		}
		filename, err := common.WriteFile(fileBytes, base)
		if err != nil {
			return err
		}

		user.Image = filename
		e := conn.Insert(user)
		if e != nil {
			return echo.NewHTTPError(e.Status, e.Error)
		}
		return nil
	}
}

func ParseRequestAndWriteFile(ctx echo.Context) ([]byte, error) {
	ctx.Request().Body = http.MaxBytesReader(ctx.Response(), ctx.Request().Body, images.MaxUploadSize)
	if err := ctx.Request().ParseMultipartForm(images.MaxUploadSize); err != nil {
		return []byte{}, echo.NewHTTPError(400, err.Error())
	}
	defer ctx.Request().Body.Close()

	file, _, err := ctx.Request().FormFile("file")
	if err != nil {
		return []byte{}, echo.NewHTTPError(400, err.Error())
	}
	defer func() {
		err := file.Close()
		if err != nil {
			ctx.Logger().Warn("can not close file")
		}
	}()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return []byte{}, echo.NewHTTPError(400, err.Error())
	}
	return fileBytes, nil
}
