package handlers

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/auth"

	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/inmemory"
	"github.com/gorilla/mux"
)

func InitDatabase() *inmemory.DB {
	db := inmemory.Init()
	db.FakeFillDB()
	return db
}

func CheckStatusCodeAndResponse(t *testing.T, caseNumber int, w *httptest.ResponseRecorder, expectedCode int, expectedRespPart string) {
	if w.Code != expectedCode {
		t.Errorf("[%d] wrong StatusCode: got %d, expected %d", caseNumber, w.Code, expectedCode)
	}

	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("[%d] an error reading response body occurred", caseNumber)
		return
	}
	defer resp.Body.Close()

	bodyStr := string(body)
	if !strings.Contains(string(body), expectedRespPart) {
		t.Errorf("[%d] wrong ResponsePart: got %+v\nExpected part of it: %+v", caseNumber, bodyStr, expectedRespPart)
	}
}

type TestCase struct {
	ID           string
	Cookie       http.Cookie
	Body         io.Reader
	ResponsePart string
	StatusCode   int
}

func TestGetFilm(t *testing.T) {
	db := InitDatabase()

	cases := []TestCase{
		{
			ID:           "1",
			ResponsePart: `"title":"Матрица"`,
			StatusCode:   http.StatusOK,
		},
		{
			ID:           "500",
			ResponsePart: `"error":"no film with the id: 500"`,
			StatusCode:   http.StatusNotFound,
		},
	}

	for caseNum, item := range cases {
		url := "http://167.71.5.55/api/films/" + item.ID
		req := httptest.NewRequest("GET", url, item.Body)
		w := httptest.NewRecorder()

		// Need to create a router that we can pass the request through so that the vars will be added to the context
		router := mux.NewRouter()
		router.HandleFunc("/api/films/{film_id:[0-9]+}", GetHandlerFilm(db))
		router.ServeHTTP(w, req)

		CheckStatusCodeAndResponse(t, caseNum, w, item.StatusCode, item.ResponsePart)
	}
}

func TestGetUser(t *testing.T) {
	db := InitDatabase()

	cases := []TestCase{
		{
			ID:           "1",
			ResponsePart: `"username":"Anton"`,
			StatusCode:   http.StatusOK,
		},
		{
			ID:           "500",
			ResponsePart: `"error":"no user with id: 500"`,
			StatusCode:   http.StatusNotFound,
		},
		{
			ID:           "5a",
			ResponsePart: `404 page not found`,
			StatusCode:   http.StatusNotFound,
		},
	}

	for caseNum, item := range cases {
		url := "http://167.71.5.55/api/users/" + item.ID
		req := httptest.NewRequest("GET", url, item.Body)
		w := httptest.NewRecorder()

		// Need to create a router that we can pass the request through so that the vars will be added to the context
		router := mux.NewRouter()
		router.HandleFunc("/api/users/{user_id:[0-9]+}", GetHandlerUser(db))
		router.ServeHTTP(w, req)

		CheckStatusCodeAndResponse(t, caseNum, w, item.StatusCode, item.ResponsePart)
	}
}

func TestGetUsers(t *testing.T) {
	db := InitDatabase()
	cookie := auth.GenerateCookie(configs.CookieAuthName, "cookie")
	db.InsertCookie(cookie, 0)

	cases := []TestCase{
		{
			ResponsePart: `"error":"no cookie"`,
			StatusCode:   http.StatusUnauthorized,
		},
		{
			Cookie:       http.Cookie{Name: configs.CookieAuthName, Value: "fake"},
			ResponsePart: `"error":"invalid cookie"`,
			StatusCode:   http.StatusUnauthorized,
		},
		{
			Cookie:       cookie,
			ResponsePart: `"id":0,"username":"Oleg"`,
			StatusCode:   http.StatusOK,
		},
	}

	for caseNum, item := range cases {
		url := "http://167.71.5.55/api/users/"
		req := httptest.NewRequest("GET", url, item.Body)
		req.AddCookie(&item.Cookie)
		w := httptest.NewRecorder()

		handler := GetHandlerUsers(db)
		handler(w, req)

		CheckStatusCodeAndResponse(t, caseNum, w, item.StatusCode, item.ResponsePart)
	}
}

func TestGetUsersCreate(t *testing.T) {
	db := InitDatabase()
	cookie := auth.GenerateCookie(configs.CookieAuthName, "cookie")
	db.InsertCookie(cookie, 0)

	cases := []TestCase{
		{
			ResponsePart: `"error":"EOF"`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Cookie:       http.Cookie{Name: configs.CookieAuthName, Value: "fake"},
			ResponsePart: `"error":"EOF"`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Cookie:       cookie,
			ResponsePart: `"error":"user is already logged in"`,
			StatusCode:   http.StatusForbidden,
		},
		{
			Body:         strings.NewReader(`{"email":"oleg@mail.ru"}`),
			ResponsePart: `"error":"user with the email already exists"`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Body:       strings.NewReader(`{"email":"katya@mail.ru","password":"qwerty!23"}`),
			StatusCode: http.StatusOK,
		},
	}

	for caseNum, item := range cases {
		url := "http://167.71.5.55/api/users/"
		req := httptest.NewRequest("POST", url, item.Body)
		req.AddCookie(&item.Cookie)
		w := httptest.NewRecorder()

		handler := GetHandlerUsersCreate(db)
		handler(w, req)

		CheckStatusCodeAndResponse(t, caseNum, w, item.StatusCode, item.ResponsePart)
	}
}

func TestGetUsersUpdate(t *testing.T) {
	db := InitDatabase()
	cookie := auth.GenerateCookie(configs.CookieAuthName, "cookie")
	db.InsertCookie(cookie, 0)

	cases := []TestCase{
		{
			ResponsePart: `"error":"EOF"`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Cookie:       http.Cookie{Name: configs.CookieAuthName, Value: "fake"},
			ResponsePart: `"error":"EOF"`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Cookie:       cookie,
			ResponsePart: `"error":"EOF"`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Cookie:       http.Cookie{Name: configs.CookieAuthName, Value: "fake"},
			Body:         strings.NewReader(`{"Username":"OLEG"}`),
			ResponsePart: `"error":"no user with the cookie"`,
			StatusCode:   http.StatusUnauthorized,
		},
		{
			Body:         strings.NewReader(`{"Username":"OLEG"}`),
			ResponsePart: `"error":"http: named cookie not present"`,
			StatusCode:   http.StatusUnauthorized,
		},
		{
			Cookie:     cookie,
			Body:       strings.NewReader(`{"username":"OLEG"}`),
			StatusCode: http.StatusOK,
		},
	}

	for caseNum, item := range cases {
		url := "http://167.71.5.55/api/users/"
		req := httptest.NewRequest("PUT", url, item.Body)
		req.AddCookie(&item.Cookie)
		w := httptest.NewRecorder()

		handler := GetHandlerUsersUpdate(db)
		handler(w, req)

		CheckStatusCodeAndResponse(t, caseNum, w, item.StatusCode, item.ResponsePart)
	}
}

func TestLoginCheck(t *testing.T) {
	db := InitDatabase()
	cookie := auth.GenerateCookie(configs.CookieAuthName, "cookie")
	db.InsertCookie(cookie, 0)

	cases := []TestCase{
		{
			StatusCode: http.StatusUnauthorized,
		},
		{
			Cookie:     http.Cookie{Name: configs.CookieAuthName, Value: "fake"},
			StatusCode: http.StatusUnauthorized,
		},
		{
			Cookie:     cookie,
			StatusCode: http.StatusOK,
		},
	}

	for caseNum, item := range cases {
		url := "http://167.71.5.55/api/sessions/"
		req := httptest.NewRequest("GET", url, item.Body)
		req.AddCookie(&item.Cookie)
		w := httptest.NewRecorder()

		handler := GetHandlerLoginCheck(db)
		handler(w, req)

		CheckStatusCodeAndResponse(t, caseNum, w, item.StatusCode, item.ResponsePart)
	}
}

func TestLogin(t *testing.T) {
	db := InitDatabase()
	cookie := auth.GenerateCookie(configs.CookieAuthName, "cookie")
	db.InsertCookie(cookie, 1)

	cases := []TestCase{
		{
			ResponsePart: `"error":"invalid json; EOF"`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Cookie:       http.Cookie{Name: configs.CookieAuthName, Value: "fake"},
			ResponsePart: `"error":"invalid json; EOF"`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Cookie:       cookie,
			ResponsePart: `"error":"invalid json; EOF"`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Cookie:       http.Cookie{Name: configs.CookieAuthName, Value: "fake"},
			Body:         strings.NewReader(`{"email":"anton@mail.ru","password":"qwe523"}`),
			ResponsePart: `"error":"invalid cookie"`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Cookie:       cookie,
			Body:         strings.NewReader(`{"email":"anton@mail.ru","password":"qwe523"}`),
			ResponsePart: `"error":"already logged in"`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Body:         strings.NewReader(`{"email":"anton@mail.ru","password":"lalala"}`),
			ResponsePart: `"error":"invalid credentials"`,
			StatusCode:   http.StatusBadRequest,
		},
		{
			Body:         strings.NewReader(`{"email":"anton@mail.ru","password":"qwe523"}`),
			ResponsePart: `"error":"invalid credentials"`,
			StatusCode:   http.StatusBadRequest, // TODO: must be ok
		},
	}

	for caseNum, item := range cases {
		url := "http://167.71.5.55/api/sessions/"
		req := httptest.NewRequest("POST", url, item.Body)
		req.AddCookie(&item.Cookie)
		w := httptest.NewRecorder()

		handler := GetHandlerLogin(db)
		handler(w, req)

		CheckStatusCodeAndResponse(t, caseNum, w, item.StatusCode, item.ResponsePart)
	}
}
