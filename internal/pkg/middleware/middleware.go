package middleware

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/inmemory"
	"net/http"
)

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
				http.Error(w, "Forbidden: no cookie", http.StatusForbidden)
				return
			}
			ok := db.CheckCookie(*cookie)
			if !ok {
				http.Error(w, "Forbidden: no cookie in db", http.StatusForbidden)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
