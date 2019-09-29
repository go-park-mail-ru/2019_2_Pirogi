package auth

import (
	"errors"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	Error "github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/error"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/inmemory"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/user"
	"net/http"
	"time"
)

func GenerateAuthCookie(value string) http.Cookie {
	expiration := time.Now().Add(configs.CookieAuthDuration)
	cookie := http.Cookie{
		Name:     configs.CookieAuthName,
		Value:    user.GetMD5Hash(value),
		Expires:  expiration,
		HttpOnly: true,
		Path:     "/",
	}
	return cookie
}

func GenerateDeAuthCookie() http.Cookie {
	cookie := http.Cookie{
		Name:     configs.CookieAuthName,
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	}
	return cookie
}

func Login(w http.ResponseWriter, r *http.Request, db *inmemory.DB, email, password string) *models.Error {
	_, err := r.Cookie(configs.CookieAuthName)
	isAuth := err != http.ErrNoCookie
	if !isAuth {
		u, ok := db.FindByEmail(email)
		if !ok {
			return Error.New(404, "no user with this email")
		}
		if u.Password != password {
			return Error.New(400, "incorrect password")
		}
		cookie := GenerateAuthCookie(email)
		e := db.Insert(cookie, u.ID)

		if e != nil {
			return e
		}
		http.SetCookie(w, &cookie)
		return nil
	}
	return Error.New(400, "already logged in")
}

func LoginCheck(w http.ResponseWriter, r *http.Request, db *inmemory.DB) bool {
	session, err := r.Cookie(configs.CookieAuthName)
	if err != nil {
		return false
	}
	_, ok := db.FindUserByCookie(*session)
	return ok
}

func Logout(w http.ResponseWriter, r *http.Request, db *inmemory.DB) error {
	session, err := r.Cookie(configs.CookieAuthName)
	isAuth := err != http.ErrNoCookie
	if !isAuth {
		return errors.New("user is not authorized")
	}
	if err == nil {

		cookie := GenerateDeAuthCookie()
		http.SetCookie(w, &cookie)
		db.Delete(*session)
	}
	return err
}
