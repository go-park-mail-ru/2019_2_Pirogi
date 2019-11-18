package usecases

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/domains"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/error"
	"github.com/labstack/echo"
	"net/http"
)

type AuthInteractor struct {
	UserRepository   domains.UserRepository
	CookieRepository domains.CookieRepository
}

func (interactor *AuthInteractor) GetUserByRequest(r *http.Request) (domains.User, bool) {
	session, err := interactor.CookieRepository.GetFromRequest(r, configs.Default.CookieAuthName)
	if err != nil {
		return domains.User{}, false
	}
	user, ok := interactor.UserRepository.GetByCookie(*session)
	if !ok {
		return domains.User{}, false
	}
	return user, true
}

func (interactor *AuthInteractor) Login(ctx echo.Context, email, password string) *domains.Error {
	_, err := interactor.CookieRepository.GetFromRequest(ctx.Request(), configs.Default.CookieAuthName)
	if err == nil {
		return error.New(400, "already logged in")
	}
	u, ok := interactor.UserRepository.GetByEmail(email)
	if !ok || u.CheckPassword(password) {
		return error.New(400, "invalid credentials")
	}
	cookie := domains.Cookie{}
	cookie.Generate(configs.Default.CookieAuthName, email)
	_, e := interactor.CookieRepository.Insert(cookie)
	if e != nil {
		return error.New(500, e.Error())
	}
	interactor.CookieRepository.SetOnResponse(ctx.Response(), &cookie)
	return nil
}

func (interactor *AuthInteractor) LoginCheck(ctx echo.Context) bool {
	session, err := interactor.CookieRepository.GetFromRequest(ctx.Request(), configs.Default.CookieAuthName)
	if err != nil {
		return false
	}
	_, err = interactor.CookieRepository.Insert(*session)
	if err != nil {
		return false
	}
	return true
}

func (interactor *AuthInteractor) Logout(ctx echo.Context) *domains.Error {
	session, err := interactor.CookieRepository.GetFromRequest(ctx.Request(), configs.Default.CookieAuthName)
	if err != nil {
		return error.New(http.StatusUnauthorized, "user is not authorized")
	}
	interactor.CookieRepository.SetOnResponse(ctx.Response(), session)
	interactor.CookieRepository.Delete(interactor.CookieRepository.Find(session))
	return nil
}
