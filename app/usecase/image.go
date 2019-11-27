package usecase

import (
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/repository"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/files"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/hash"
	"github.com/labstack/echo"
)

type ImageUsecase interface {
	GetUserByContext(ctx echo.Context) (model.User, *model.Error)
	UpdateUserImage(user model.User, filename string) *model.Error
	GenerateFilename(ctx echo.Context, body []byte) (fullPath, filename string, err *model.Error)
}

type imageUsecase struct {
	cookieRepo repository.CookieRepository
	userRepo   repository.UserRepository
}

func NewImageUsecase(cookieRepo repository.CookieRepository, userRepo repository.UserRepository) *imageUsecase {
	return &imageUsecase{
		cookieRepo: cookieRepo,
		userRepo:   userRepo,
	}
}

func (u *imageUsecase) GetUserByContext(ctx echo.Context) (model.User, *model.Error) {
	return u.cookieRepo.GetUserByContext(ctx)
}

func (u *imageUsecase) UpdateUserImage(user model.User, filename string) *model.Error {
	user.Image = model.Image(filename)
	err := u.userRepo.Update(user)
	return err
}

func (u *imageUsecase) GenerateFilename(ctx echo.Context, body []byte) (fullPath, filename string, err *model.Error) {
	var base string
	switch {
	case strings.Contains(ctx.Request().URL.Path, "users"):
		base = configs.Default.UsersImageUploadPath
	case strings.Contains(ctx.Request().URL.Path, "films"):
		base = configs.Default.FilmsImageUploadPath
	case strings.Contains(ctx.Request().URL.Path, "persons"):
		base = configs.Default.PersonsImageUploadPath
	default:
		return "", "", model.NewError(http.StatusBadRequest, "wrong path")
	}
	ending, err := files.DetectContentType(body)
	if err != nil {
		return "", "", err
	}
	filename = hash.SHA1(path.Join(time.Now().String())) + ending
	fullPath = path.Join(base, filename)
	return fullPath, filename, nil
}
