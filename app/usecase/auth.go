package usecase

import (
	"context"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/repository"
	v1 "github.com/go-park-mail-ru/2019_2_Pirogi/app/infrastructure/microservices/sessions/protobuf"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/network"
	"github.com/labstack/echo"
	"net/http"
)

type AuthUsecase interface {
	Login(ctx echo.Context, email, password string) (int, *model.Error)
	LoginCheck(ctx echo.Context) (int, bool)
	Logout(ctx echo.Context) *model.Error
}

type authUsecase struct {
	subscriptionRepo repository.SubscriptionRepository
	rpcClient        v1.AuthServiceClient
}

func NewAuthUsecase(subscriptionRepo repository.SubscriptionRepository, rpcClient v1.AuthServiceClient) *authUsecase {
	return &authUsecase{
		subscriptionRepo: subscriptionRepo,
		rpcClient:        rpcClient,
	}
}

func (u *authUsecase) Login(ctx echo.Context, email, password string) (int, *model.Error) {
	_, err := network.GetCookieFromContext(ctx, configs.Default.CookieAuthName)
	if err == nil {
		return -1, model.NewError(400, "Пользователь уже авторизован")
	}
	response, e := u.rpcClient.Login(context.Background(), &v1.LoginRequest{
		Email:    email,
		Password: password,
	})
	if e != nil {
		return -1, model.NewError(500, e.Error())
	}
	cookie := model.Cookie{}
	cookie.GenerateAuthCookie(model.ID(response.UserID), configs.Default.CookieAuthName, response.CookieValue)
	network.SetCookieOnContext(&ctx, cookie)
	subscription, err := u.subscriptionRepo.Find(model.ID(response.UserID))
	var newEventsNumber int
	for _, event := range subscription.SubscriptionEvents {
		if !event.IsRead {
			newEventsNumber++
		}
	}
	return newEventsNumber, nil
}

func (u *authUsecase) LoginCheck(ctx echo.Context) (int, bool) {
	cookie, err := network.GetCookieFromContext(ctx, configs.Default.CookieAuthName)
	if err != nil {
		return -1, false
	}
	response, e := u.rpcClient.LoginCheck(context.Background(), &v1.LoginCheckRequest{
		CookieValue: cookie.Cookie.Value,
	})
	if e != nil {
		return -1, false
	}
	subscription, err := u.subscriptionRepo.Find(model.ID(response.UserID))
	var newEventsNumber int
	for _, event := range subscription.SubscriptionEvents {
		if !event.IsRead {
			newEventsNumber++
		}
	}
	return newEventsNumber, true
}

func (u *authUsecase) Logout(ctx echo.Context) *model.Error {
	session, err := network.GetCookieFromContext(ctx, configs.Default.CookieAuthName)
	if err != nil {
		return model.NewError(http.StatusUnauthorized, "Пользователь не авторизован")
	}
	_, e := u.rpcClient.Logout(context.Background(), &v1.LogoutRequest{
		CookieValue: session.Cookie.Value,
	})
	if e != nil {
		return model.NewError(500, e.Error())
	}
	session.Expire()
	network.SetCookieOnContext(&ctx, session)
	return nil
}
