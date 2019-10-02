package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	error "github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/error"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/inmemory"
)

func LoggingMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: понять как не открывать файл каждый раз и проверять его наличие
		if f, err := os.OpenFile(configs.AccessLog, os.O_APPEND|os.O_WRONLY, os.ModeAppend); err != nil {
			log.Fatal("Can not open file to log: ", err.Error())
		} else {
			_, _ = fmt.Fprintf(f, "%s %s %s %s \n", time.Now().Format("02/01 15:04:05"), r.Method, r.URL, r.Host)
		}
		next.ServeHTTP(w, r)
	})
}

func HeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
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
			if r.Method == http.MethodGet || r.Method == http.MethodPost && (r.URL.Path == "/api/users/" || r.URL.Path == "/api/sessions/") {
				next.ServeHTTP(w, r)
				return
			}

			cookie, err := r.Cookie(configs.CookieAuthName)
			if err != nil {
				error.Render(w, error.New(401, "no cookie"))
				return
			}
			ok := db.CheckCookie(*cookie)
			if !ok {
				error.Render(w, error.New(401, "no cookie in db"))
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
