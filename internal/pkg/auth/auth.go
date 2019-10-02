package auth

import (
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	error "github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/error"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/inmemory"
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

func ExpireCookie(cookie *http.Cookie) {
	// TODO: понять, почему кука просрачивается только при таких параметрах
	cookie.Expires = time.Unix(0, 0)
	cookie.Path = "/"
	cookie.HttpOnly = true
}

func Login(w http.ResponseWriter, r *http.Request, db *inmemory.DB, email, password string) *models.Error {
	cookie, err := r.Cookie(configs.CookieAuthName)
	if err != nil {
		u, ok := db.FindByEmail(email)
		if !ok || u.Password != password {
			return error.New(400, "invalid credentials")
		}
		cookie := GenerateCookie("cinsear_session", email)
		e := db.InsertCookie(cookie, u.ID)
		if e != nil {
			return e
		}
		http.SetCookie(w, &cookie)
		return nil
	}
	if cookie != nil {
		if _, ok := db.FindUserByCookie(*cookie); !ok {
			return error.New(400, "invalid cookie")
		}
	}
	return error.New(400, "already logged in")
}

func LoginCheck(_ http.ResponseWriter, r *http.Request, db *inmemory.DB) bool {
	session, err := r.Cookie(configs.CookieAuthName)
	if err != nil {
		return false
	}
	_, ok := db.FindUserByCookie(*session)
	return ok
}

func Logout(w http.ResponseWriter, r *http.Request, db *inmemory.DB) *models.Error {
	session, err := r.Cookie(configs.CookieAuthName)
	if err != nil {
		return error.New(401, "user is not authorized")
	}
	ExpireCookie(session)
	http.SetCookie(w, session)
	db.DeleteCookie(*session)
	return nil
}
