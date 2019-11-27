package sessions

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/repository"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/infrastructure/microservices/sessions/protobuf"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/hash"
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
	userRepo   repository.UserRepository
	cookieRepo repository.CookieRepository
}

func (u *authManager) Login(ctx context.Context, request *v1.LoginRequest) (*v1.LoginResponse, error) {
	user, err := u.userRepo.GetByEmail(request.Email)
	zap.S().Debug(user.Email, user.Password)
	zap.S().Debug(request.Email, request.Password)
	if err != nil || !user.IsPasswordCorrect(hash.SHA1(request.Password)) {
		return &v1.LoginResponse{}, errors.New("400, Неверная почта и/или пароль")
	}
	cookie := model.Cookie{}
	cookie.GenerateAuthCookie(user.ID, configs.Default.CookieAuthName, hash.SHA1(request.Email))
	e := u.cookieRepo.Insert(cookie)
	if e != nil {
		return &v1.LoginResponse{}, e.Common()
	}
	zap.S().Debug(cookie.Cookie.Value)
	return &v1.LoginResponse{UserID: int64(user.ID), CookieValue: cookie.Cookie.Value}, nil
}

func (u *authManager) LoginCheck(ctx context.Context, request *v1.LoginCheckRequest) (*v1.LoginCheckResponse, error) {
	zap.S().Debug(request.CookieValue)
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
	zap.S().Debug(request.CookieValue)
	zap.S().Debug(user.Email)
	if err != nil {
		return &v1.LogoutResponse{}, model.NewError(http.StatusUnauthorized, "Пользователь не авторизован").Common()
	}
	var cookie model.Cookie
	cookie.GenerateAuthCookie(user.ID, configs.Default.CookieAuthName, request.CookieValue)
	u.cookieRepo.Delete(cookie)
	zap.S().Debug("cookie deleted")
	return &v1.LogoutResponse{}, nil
}
