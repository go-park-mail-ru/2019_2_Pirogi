package http

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/usecase"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/files"
	"github.com/labstack/echo"
)

func GetImagesHandler(usecase usecase.ImageUsecase) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		user, err := usecase.GetUserByContext(ctx)
		if err != nil {
			return err.HTTP()
		}
		fileBytes, err := files.ReadFileByteFromContext(ctx)
		if err != nil {
			return err.HTTP()
		}

		fullPath, filename, err := usecase.GenerateFilename(ctx, fileBytes)
		if err != nil {
			return err.HTTP()
		}
		err = files.WriteFile(fullPath, fileBytes)
		if err != nil {
			return err.HTTP()
		}
		err = usecase.UpdateUserImage(user, filename)
		if err != nil {
			return err.HTTP()
		}
		return nil
	}
}
