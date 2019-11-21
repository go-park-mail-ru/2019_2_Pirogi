package usecase

import (
	"github.com/asaskevich/govalidator"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/repository"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/hash"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/network"
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type UserUsecase interface {
	GetUserByContext(ctx echo.Context) (model.User, *model.Error)
	GetUserTruncByteByID(id model.ID) ([]byte, *model.Error)
	CreateUserNewFromContext(ctx echo.Context) *model.Error
	UpdateUserFromContext(ctx echo.Context) *model.Error
}

func NewUserUsecase(userRepo repository.UserRepository, cookieRepo repository.CookieRepository) *userUsecase {
	return &userUsecase{
		userRepo:   userRepo,
		cookieRepo: cookieRepo,
	}
}

type userUsecase struct {
	userRepo   repository.UserRepository
	cookieRepo repository.CookieRepository
}

func (u userUsecase) GetUserByContext(ctx echo.Context) (model.User, *model.Error) {
	cookie, err := u.cookieRepo.GetFromRequest(ctx.Request(), configs.Default.CookieAuthName)
	if err != nil {
		return model.User{}, err
	}
	return u.cookieRepo.GetUserByCookie(cookie)
}

func (u userUsecase) GetUserTruncByteByID(id model.ID) ([]byte, *model.Error) {
	user, err := u.userRepo.Get(id)
	zap.S().Debug(user)
	if err != nil {
		return nil, err
	}
	userTrunc := user.Trunc()
	body, e := userTrunc.MarshalJSON()
	if e != nil {
		return nil, model.NewError(500, e.Error())
	}
	return body, nil
}

func (u userUsecase) CreateUserNewFromContext(ctx echo.Context) *model.Error {
	body, err := network.ReadBody(ctx)
	if err != nil {
		return err
	}
	var userNew model.UserNew
	e := userNew.UnmarshalJSON(body)
	if e != nil {
		return model.NewError(400, e.Error())
	}
	zap.S().Debug(userNew)
	_, err = u.userRepo.GetByEmail(userNew.Email)
	if err == nil {
		return model.NewError(400, "user with the email is already existed")
	}
	err = u.userRepo.PrepareUserNew(&userNew)
	if err != nil {
		return err
	}
	err = u.userRepo.Insert(userNew)
	if err != nil {
		return err
	}
	user, err := u.userRepo.GetByEmail(userNew.Email)
	var cookie model.Cookie
	cookie.GenerateAuthCookie(user.ID, configs.Default.CookieAuthName, hash.SHA1(userNew.Password+userNew.Email+time.Now().String()))
	err = u.cookieRepo.Insert(cookie)
	if err != nil {
		return err
	}
	u.cookieRepo.SetOnResponse(ctx.Response(), &cookie)
	return nil
}

func (u userUsecase) UpdateUserFromContext(ctx echo.Context) *model.Error {
	body, err := network.ReadBody(ctx)
	if err != nil {
		return err
	}
	var updateUser model.User
	e := updateUser.UnmarshalJSON(body)
	if e != nil {
		return model.NewError(400, e.Error())
	}
	_, e = govalidator.ValidateStruct(updateUser)
	if e != nil {
		return model.NewError(http.StatusBadRequest, e.Error())
	}
	session, err := u.cookieRepo.GetFromRequest(ctx.Request(), configs.Default.CookieAuthName)
	if err != nil {
		return err
	}
	user, err := u.cookieRepo.GetUserByCookie(session)
	if err != nil {
		return err
	}
	switch {
	case updateUser.Username != "":
		user.Username = updateUser.Username
		fallthrough
	case updateUser.Description != "":
		user.Description = updateUser.Description
	}
	zap.S().Debug(updateUser)
	zap.S().Debug(user)
	err = u.userRepo.Update(user)
	return err
}
