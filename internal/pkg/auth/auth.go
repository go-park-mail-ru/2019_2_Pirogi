package auth

import (
	"github.com/labstack/echo"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/database"
	error "github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/error"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/user"
)

// Creates cookie with value = MD5(value)
func GenerateCookie(cookieName, value string) http.Cookie {
	expiration := time.Now().Add(configs.CookieAuthDuration)
	cookie := http.Cookie{
		Name:     cookieName,
		Value:    user.GetMD5Hash(value),
		Expires:  expiration,
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
	}
	return cookie
}

func GetUserByRequest(r *http.Request, conn database.Database) (models.User, bool) {
	session, err := r.Cookie(configs.CookieAuthName)
	if err != nil {
		return models.User{}, false
	}
	foundUser, ok := conn.FindUserByCookie(session)
	if !ok {
		return models.User{}, false
	}
	return foundUser, true
}

func ExpireCookie(cookie *http.Cookie) {
	cookie.Expires = time.Unix(0, 0)
	cookie.Path = "/"
	cookie.HttpOnly = true
}

func Login(ctx echo.Context, db database.Database, email, password string) *models.Error {
	cookie, err := ctx.Request().Cookie(configs.CookieAuthName)
	if err != nil {
		u, ok := db.FindByEmail(email)
		if !ok || u.Password != password {
			return error.New(400, "invalid credentials")
		}
		cookie := GenerateCookie("cinsear_session", email)
		e := db.InsertCookie(&cookie, u.ID)
		if e != nil {
			return e
		}
		http.SetCookie(ctx.Response(), &cookie)
		return nil
	}
	if cookie != nil {
		if _, ok := db.FindUserByCookie(cookie); !ok {
			return error.New(400, "invalid cookie")
		}
	}
	return error.New(400, "already logged in")
}

func LoginCheck(_ http.ResponseWriter, r *http.Request, db database.Database) bool {
	session, err := r.Cookie(configs.CookieAuthName)
	if err != nil {
		return false
	}
	_, ok := db.FindUserByCookie(session)
	return ok
}

func Logout(ctx echo.Context, db database.Database) *models.Error {
	session, err := ctx.Request().Cookie(configs.CookieAuthName)
	if err != nil {
		return error.New(401, "user is not authorized")
	}
	ExpireCookie(session)
	http.SetCookie(ctx.Response(), session)
	db.DeleteCookie(*session)
	return nil
}
