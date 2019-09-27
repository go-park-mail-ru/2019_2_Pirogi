package auth

import (
	"errors"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/inmemory"
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

func Login(w http.ResponseWriter, r *http.Request, db *inmemory.DB, email, password string) error {
	_, err := r.Cookie(configs.CookieAuthName)
	isAuth := err != http.ErrNoCookie
	if !isAuth {
		u, ok := db.FindByEmail(email)
		if !ok {
			return errors.New("no user with this email")
		}
		if u.Password != password {
			return errors.New("incorrect password")
		}
		cookie := GenerateAuthCookie(email)
		err = db.Insert(cookie)

		if err != nil {
			return errors.New("db error while inserting cookie: " + err.Error())
		}
		http.SetCookie(w, &cookie)
		return nil
	}
	return errors.New("already logged in")
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
