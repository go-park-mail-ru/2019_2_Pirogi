package usecase

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/repository"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/labstack/echo"
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
}

func NewAuthUsecase(userRepo repository.UserRepository, cookieRepo repository.CookieRepository,
	subscriptionRepository repository.SubscriptionRepository) *authUsecase {
	return &authUsecase{
		userRepo:         userRepo,
		cookieRepo:       cookieRepo,
		subscriptionRepo: subscriptionRepository,
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
	user, err := u.userRepo.GetByEmail(email)
	if err != nil || user.CheckPassword(password) {
		return -1, model.NewError(400, "Неверная почта и/или пароль")
	}
	cookie := model.Cookie{}
	cookie.GenerateAuthCookie(user.ID, configs.Default.CookieAuthName, email)
	e := u.cookieRepo.Insert(cookie)
	if e != nil {
		return -1, e
	}
	u.cookieRepo.SetOnResponse(ctx.Response(), &cookie)
	subscription, err := u.subscriptionRepo.Find(user.ID)
	var newEventsNumber int
	for _, event := range subscription.SubscriptionEvents {
		if !event.IsRead {
			newEventsNumber++
		}
	}
	return newEventsNumber, nil
}

func (u *authUsecase) LoginCheck(ctx echo.Context) (int, bool) {
	user, err := u.cookieRepo.GetUserByContext(ctx)
	if err != nil {
		return -1, false
	}
	subscription, err := u.subscriptionRepo.Find(user.ID)
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
	session.Expire()
	u.cookieRepo.SetOnResponse(ctx.Response(), &session)
	u.cookieRepo.Delete(session)
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
