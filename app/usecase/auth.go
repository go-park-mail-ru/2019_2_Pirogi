package usecase

import (
	"context"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/repository"
	v1 "github.com/go-park-mail-ru/2019_2_Pirogi/cmd/sessions/protobuf"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net/http"
)

type AuthUsecase interface {
	Login(ctx echo.Context, email, password string) (int, *model.Error)
	LoginCheck(ctx echo.Context) (int, bool)
	Logout(ctx echo.Context) *model.Error
	CheckCookieExisting(ctx echo.Context, cookieName string) bool
	GetCookie(ctx echo.Context, cookieName string) (model.Cookie, *model.Error)
}

type authUsecase struct {
	userRepo         repository.UserRepository
	cookieRepo       repository.CookieRepository
	subscriptionRepo repository.SubscriptionRepository
	rpcClient        v1.AuthServiceClient
}

func NewAuthUsecase(userRepo repository.UserRepository, cookieRepo repository.CookieRepository,
	subscriptionRepository repository.SubscriptionRepository) *authUsecase {
	grcpConn, err := grpc.Dial(
		"sessions:8081",
		grpc.WithInsecure(),
	)
	if err != nil {
		zap.S().Error(err.Error())
	}

	client := v1.NewAuthServiceClient(grcpConn)
	return &authUsecase{
		userRepo:         userRepo,
		cookieRepo:       cookieRepo,
		subscriptionRepo: subscriptionRepository,
		rpcClient:        client,
	}
}
func (u *authUsecase) CheckCookieExisting(ctx echo.Context, cookieName string) bool {
	_, err := ctx.Cookie(cookieName)
	if err != nil {
		return false
	}
	return true
}

func (u *authUsecase) Login(ctx echo.Context, email, password string) (int, *model.Error) {
	_, err := u.cookieRepo.GetCookieFromRequest(ctx.Request(), configs.Default.CookieAuthName)
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
	u.cookieRepo.SetOnResponse(ctx.Response(), &cookie)
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
	cookie, err := u.cookieRepo.GetCookieFromRequest(ctx.Request(), configs.Default.CookieAuthName)
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
	session, err := u.cookieRepo.GetCookieFromRequest(ctx.Request(), configs.Default.CookieAuthName)
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
	u.cookieRepo.SetOnResponse(ctx.Response(), &session)
	return nil
}

func (u *authUsecase) GetCookie(ctx echo.Context, cookieName string) (model.Cookie, *model.Error) {
	if !u.CheckCookieExisting(ctx, cookieName) {
		return model.Cookie{}, model.NewError(400, "Пользователь не авторизован")
	}
	var cookie model.Cookie
	cookieCommon, _ := ctx.Cookie(cookieName)
	cookie.CopyFromCommon(cookieCommon)
	return cookie, nil
}
