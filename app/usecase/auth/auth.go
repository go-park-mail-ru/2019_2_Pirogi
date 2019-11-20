package usecase

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/repository"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/service"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/labstack/echo"
	"net/http"
)

type AuthUsecase interface {
	GetUserByRequest(r *http.Request) (model.User, bool)
	Login(ctx echo.Context, email, password string) *model.Error
	LoginCheck(ctx echo.Context) bool
	Logout(ctx echo.Context) *model.Error
}

type authUsecase struct {
	userRepo      repository.UserRepository
	cookieRepo    repository.CookieRepository
	userService   *service.UserService
	cookieService *service.CookieService
}

func NewAuthUsecase(userRepo repository.UserRepository, cookieRepo repository.CookieRepository,
	userService *service.UserService, cookieService *service.CookieService) *authUsecase {
	return &authUsecase{
		userRepo:      userRepo,
		cookieRepo:    cookieRepo,
		userService:   userService,
		cookieService: cookieService,
	}
}

func (u *authUsecase) GetUserByRequest(r *http.Request) (model.User, bool) {
	session, err := u.cookieRepo.GetFromRequest(r, configs.Default.CookieAuthName)
	if err != nil {
		return model.User{}, false
	}
	user, ok := u.userRepo.GetByCookie(*session)
	if !ok {
		return model.User{}, false
	}
	return user, true
}

func (u *authUsecase) Login(ctx echo.Context, email, password string) *model.Error {
	_, err := u.cookieRepo.GetFromRequest(ctx.Request(), configs.Default.CookieAuthName)
	if err == nil {
		return model.NewError(400, "already logged in")
	}
	user, ok := u.userRepo.GetByEmail(email)
	if !ok || user.CheckPassword(password) {
		return model.NewError(400, "invalid credentials")
	}
	cookie := model.Cookie{}
	cookie.Generate(configs.Default.CookieAuthName, email)
	_, e := u.cookieRepo.Insert(cookie)
	if e != nil {
		return model.NewError(500, e.Error())
	}
	u.cookieRepo.SetOnResponse(ctx.Response(), &cookie)
	return nil
}

func (u *authUsecase) LoginCheck(ctx echo.Context) bool {
	session, err := u.cookieRepo.GetFromRequest(ctx.Request(), configs.Default.CookieAuthName)
	if err != nil {
		return false
	}
	_, err = u.cookieRepo.Insert(*session)
	if err != nil {
		return false
	}
	return true
}

func (u *authUsecase) Logout(ctx echo.Context) *model.Error {
	session, err := u.cookieRepo.GetFromRequest(ctx.Request(), configs.Default.CookieAuthName)
	if err != nil {
		return model.NewError(http.StatusUnauthorized, "user is not authorized")
	}
	u.cookieRepo.SetOnResponse(ctx.Response(), session)
	u.cookieRepo.Delete(u.cookieRepo.Find(session))
	return nil
}
