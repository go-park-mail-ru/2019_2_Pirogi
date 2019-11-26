package usecase

import (
	"context"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	v12 "github.com/go-park-mail-ru/2019_2_Pirogi/app/infrastructure/microservices/sessions/protobuf"
	v1 "github.com/go-park-mail-ru/2019_2_Pirogi/app/infrastructure/microservices/users/protobuf"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/hash"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/modelWorker"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/network"
	"github.com/labstack/echo"
	"time"
)

type UserUsecase interface {
	GetUserByContext(ctx echo.Context) (model.User, *model.Error)
	GetUserTruncByteByID(id model.ID) ([]byte, *model.Error)
	CreateUserNewFromContext(ctx echo.Context) *model.Error
	UpdateUserFromContext(ctx echo.Context) *model.Error
}

func NewUserUsecase(usersRpcClient v1.UserServiceClient, sessionsRpcClient v12.AuthServiceClient) *userUsecase {
	return &userUsecase{
		usersRpcClient:    usersRpcClient,
		sessionsRpcClient: sessionsRpcClient,
	}
}

type userUsecase struct {
	usersRpcClient    v1.UserServiceClient
	sessionsRpcClient v12.AuthServiceClient
}

func (u userUsecase) GetUserByContext(ctx echo.Context) (model.User, *model.Error) {
	cookie, err := network.GetCookieFromContext(ctx, configs.Default.CookieAuthName)
	if err != nil {
		return model.User{}, err
	}
	user, e := u.usersRpcClient.GetByCookieValue(context.Background(), &v1.CookieValue{
		CookieValue: cookie.Cookie.Value,
	})
	if e != nil {
		return model.User{}, model.NewError(404, e.Error())
	}
	var userModel model.User
	userModel.FromProtobuf(*user)
	return userModel, nil
}

func (u userUsecase) GetUserTruncByteByID(id model.ID) ([]byte, *model.Error) {
	user, e := u.usersRpcClient.GetByID(context.Background(), &v1.ID{
		ID: int64(id),
	})
	if e != nil {
		return nil, model.NewError(404, e.Error())
	}
	var userModel model.User
	userModel.FromProtobuf(*user)
	userTrunc := userModel.Trunc()
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
	_, e = u.usersRpcClient.GetByEmail(context.Background(), &v1.Email{Email: userNew.Email})
	if e == nil {
		return model.NewError(400, "user with the email is already existed")
	}
	err = modelWorker.PrepareUserNew(&userNew)
	if err != nil {
		return err
	}
	_, e = u.usersRpcClient.Create(context.Background(), &v1.UserNew{
		Email:    userNew.Email,
		Password: userNew.Password,
	})
	if e != nil {
		return model.NewError(500, e.Error())
	}
	user, e := u.usersRpcClient.GetByEmail(context.Background(), &v1.Email{
		Email: userNew.Email,
	})
	if e != nil {
		return model.NewError(404, e.Error())
	}
	_, e = u.sessionsRpcClient.Login(context.Background(), &v12.LoginRequest{
		Email:    user.Email,
		Password: user.Password,
	})
	if e != nil {
		return model.NewError(500, e.Error())
	}
	var cookie model.Cookie
	cookie.GenerateAuthCookie(model.ID(user.ID), configs.Default.CookieAuthName,
		hash.SHA1(userNew.Password+userNew.Email+time.Now().String()))
	network.SetCookieOnContext(&ctx, cookie)
	return nil
}

func (u userUsecase) UpdateUserFromContext(ctx echo.Context) *model.Error {
	cookie, err := network.GetCookieFromContext(ctx, configs.Default.CookieAuthName)
	if err != nil {
		return err
	}
	body, err := network.ReadBody(ctx)
	if err != nil {
		return err
	}
	var updateUser model.User
	e := updateUser.UnmarshalJSON(body)
	if e != nil {
		return model.NewError(400, e.Error())
	}
	err = modelWorker.PrepareUserUpdate(&updateUser)
	if err != nil {
		return err
	}
	user, e := u.usersRpcClient.GetByCookieValue(context.Background(), &v1.CookieValue{
		CookieValue: cookie.Cookie.Value,
	})
	if e != nil {
		return model.NewError(404, e.Error())
	}
	switch {
	case updateUser.Username != "":
		user.Username = updateUser.Username
		fallthrough
	case updateUser.Description != "":
		user.Description = updateUser.Description
	}
	_, e = u.usersRpcClient.Update(context.Background(), user)
	return err
}
