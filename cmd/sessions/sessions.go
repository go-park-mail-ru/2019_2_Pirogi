package main

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/repository"
	v1 "github.com/go-park-mail-ru/2019_2_Pirogi/cmd/sessions/protobuf"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"go.uber.org/zap"
	"net/http"
)

func NewAuthManager(userRepo repository.UserRepository, cookieRepo repository.CookieRepository) *authManager {
	return &authManager{
		userRepo:   userRepo,
		cookieRepo: cookieRepo,
	}
}

type authManager struct {
	userRepo         repository.UserRepository
	cookieRepo       repository.CookieRepository
	subscriptionRepo repository.SubscriptionRepository
}

func (u *authManager) Login(ctx context.Context, request *v1.LoginRequest) (*v1.LoginResponse, error) {
	user, err := u.userRepo.GetByEmail(request.Email)
	if err != nil || user.CheckPassword(request.Password) {
		return &v1.LoginResponse{}, errors.New("400, Неверная почта и/или пароль")
	}
	cookie := model.Cookie{}
	cookie.GenerateAuthCookie(user.ID, configs.Default.CookieAuthName, request.Email)
	e := u.cookieRepo.Insert(cookie)
	if e != nil {
		return &v1.LoginResponse{}, e.Common()
	}
	return &v1.LoginResponse{UserID: int64(user.ID), CookieValue: cookie.Cookie.Value}, nil
}

func (u *authManager) LoginCheck(ctx context.Context, request *v1.LoginCheckRequest) (*v1.LoginCheckResponse, error) {
	user, err := u.cookieRepo.GetUserByCookieValue(request.CookieValue)
	zap.S().Debug(user)
	if err != nil {
		return &v1.LoginCheckResponse{}, err.Common()
	}
	return &v1.LoginCheckResponse{
		UserID: int64(user.ID),
	}, nil
}

func (u *authManager) Logout(ctx context.Context, request *v1.LogoutRequest) (*v1.LogoutResponse, error) {
	user, err := u.cookieRepo.GetUserByCookieValue(request.CookieValue)
	if err != nil {
		return &v1.LogoutResponse{}, model.NewError(http.StatusUnauthorized, "Пользователь не авторизован").Common()
	}
	var cookie model.Cookie
	cookie.GenerateAuthCookie(user.ID, configs.Default.CookieAuthName, request.CookieValue)
	u.cookieRepo.Delete(cookie)
	return &v1.LogoutResponse{}, nil
}
