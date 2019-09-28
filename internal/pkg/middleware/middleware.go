package middleware

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	Error "github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/error"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/inmemory"
	"log"
	"net/http"
	"os"
)

func LoggingMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if f, err := os.OpenFile(configs.AccessLogPath+"access_log.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend); err != nil {
			log.Fatal("Can not open file to log: ", err.Error())
		} else {
			_, _ = fmt.Fprint(f, r.Method, " ", r.URL, " ", r.Host)
		}
		next.ServeHTTP(w, r)
	})
}

func HeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, post-check=0, pre-check=0")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Header().Set("Vary", "Accept-Encoding")
		next.ServeHTTP(w, r)
	})
}

func GetCheckAuthMiddleware(db *inmemory.DB) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// POST разрешен для анонимов разрешен только для регистрации
			if r.Method == http.MethodGet || r.Method == http.MethodPost && (r.URL.Path == "/api/users/" || r.URL.Path == "/api/login/") {
				next.ServeHTTP(w, r)
				return
			}

			cookie, err := r.Cookie(configs.CookieAuthName)
			if err != nil {
				Error.Render(w, Error.New(401, "no cookie"))
				return
			}
			ok := db.CheckCookie(*cookie)
			if !ok {
				Error.Render(w, Error.New(401, "no cookie in db"))
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
