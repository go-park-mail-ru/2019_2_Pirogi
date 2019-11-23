package usecase

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/repository"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"net/http"
)

type AuthUsecase interface {
	Login(ctx echo.Context, email, password string) *model.Error
	LoginCheck(ctx echo.Context) bool
	Logout(ctx echo.Context) *model.Error
	CheckCookieExisting(ctx echo.Context, cookieName string) bool
	GetCookie(ctx echo.Context, cookieName string) (model.Cookie, *model.Error)
}

type authUsecase struct {
	userRepo   repository.UserRepository
	cookieRepo repository.CookieRepository
}

func NewAuthUsecase(userRepo repository.UserRepository, cookieRepo repository.CookieRepository) *authUsecase {
	return &authUsecase{
		userRepo:   userRepo,
		cookieRepo: cookieRepo,
	}
}
func (u *authUsecase) CheckCookieExisting(ctx echo.Context, cookieName string) bool {
	_, err := ctx.Cookie(cookieName)
	if err != nil {
		return false
	}
	return true
}

func (u *authUsecase) Login(ctx echo.Context, email, password string) *model.Error {
	_, err := u.cookieRepo.GetCookieFromRequest(ctx.Request(), configs.Default.CookieAuthName)
	if err == nil {
		return model.NewError(400, "already logged in")
	}
	user, err := u.userRepo.GetByEmail(email)
	if err != nil || user.CheckPassword(password) {
		return model.NewError(400, "invalid credentials")
	}
	cookie := model.Cookie{}
	cookie.GenerateAuthCookie(user.ID, configs.Default.CookieAuthName, email)
	e := u.cookieRepo.Insert(cookie)
	if e != nil {
		return e
	}
	u.cookieRepo.SetOnResponse(ctx.Response(), &cookie)
	return nil
}

func (u *authUsecase) LoginCheck(ctx echo.Context) bool {
	session, err := u.cookieRepo.GetCookieFromRequest(ctx.Request(), configs.Default.CookieAuthName)
	if err != nil {
		return false
	}
	err = u.cookieRepo.Insert(session)
	if err != nil {
		return false
	}
	return true
}

func (u *authUsecase) Logout(ctx echo.Context) *model.Error {
	session, err := u.cookieRepo.GetCookieFromRequest(ctx.Request(), configs.Default.CookieAuthName)
	zap.S().Debug(configs.Default.CookieAuthName)
	if err != nil {
		return model.NewError(http.StatusUnauthorized, "user is not authorized")
	}
	session.Expire()
	u.cookieRepo.SetOnResponse(ctx.Response(), &session)
	u.cookieRepo.Delete(session)
	return nil
}

func (u *authUsecase) GetCookie(ctx echo.Context, cookieName string) (model.Cookie, *model.Error) {
	if !u.CheckCookieExisting(ctx, cookieName) {
		return model.Cookie{}, model.NewError(400, "no cookie")
	}
	var cookie model.Cookie
	cookieCommon, _ := ctx.Cookie(cookieName)
	cookie.CopyFromCommon(cookieCommon)
	return cookie, nil
}
