package auth

import (
	"errors"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/inmemory"
)

func Auth(w http.ResponseWriter, r *http.Request, db *inmemory.DB, email, password string) error {
	_, err := r.Cookie("session_id")
	isAuth := err != http.ErrNoCookie
	if !isAuth {
		u, ok := db.FindByEmail(email)
		if !ok {
			return errors.New("no user with this email")
		}
		if u.Password != password {
			return errors.New("incorrect password")
		}
		expiration := time.Now().Add(1 * time.Hour)
		cookie := &http.Cookie{
			Name:    "session_id",
			Value:   email,
			Expires: expiration,
		}
		http.SetCookie(w, cookie)
	}
	return nil
}

func Logout(w http.ResponseWriter, r *http.Request, db *inmemory.DB) error {
	session, err := r.Cookie("session_id")
	if err == nil {
		session.Expires = time.Now().AddDate(0, 0, -1)
		http.SetCookie(w, session)
		db.DeleteCookie(session)
	}

	isAuth := err != http.ErrNoCookie
	if !isAuth {
		w.WriteHeader(http.StatusForbidden)
	}

	return err
}
