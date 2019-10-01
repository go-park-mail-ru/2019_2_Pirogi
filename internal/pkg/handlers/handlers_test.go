package handlers

import (
    "github.com/go-park-mail-ru/2019_2_Pirogi/configs"
    "github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/inmemory"
    "github.com/gorilla/mux"
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
)

func InitDatabase() *inmemory.DB  {
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

type TestCaseGetById struct {
    ID           string
    ResponsePart string
    StatusCode   int
}

func TestGetFilm(t *testing.T) {
    db := InitDatabase()

    cases := []TestCaseGetById{
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
        req := httptest.NewRequest("GET", url, nil)
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

    cases := []TestCaseGetById{
        {
            ID:           "1",
            ResponsePart: `"user_id":1`,
            StatusCode:   http.StatusOK,
        },
        {
            ID:           "500",
            ResponsePart: `"error":"no user with id: 500"`,
            StatusCode:   http.StatusNotFound,
        },
    }

    for caseNum, item := range cases {
        url := "http://167.71.5.55/api/users/" + item.ID
        req := httptest.NewRequest("GET", url, nil)
        w := httptest.NewRecorder()

        // Need to create a router that we can pass the request through so that the vars will be added to the context
        router := mux.NewRouter()
        router.HandleFunc("/api/users/{user_id:[0-9]+}", GetHandlerUser(db))
        router.ServeHTTP(w, req)

        CheckStatusCodeAndResponse(t, caseNum, w, item.StatusCode, item.ResponsePart)
    }
}

type TestCaseGetUsers struct {
    Cookie       http.Cookie
    ResponsePart string
    StatusCode   int
}

func TestGetUsers(t *testing.T) {
    db := InitDatabase()
    cookie := http.Cookie{Name: configs.CookieAuthName, Value: "cookie"}
    db.Insert(cookie, 1)

    cases := []TestCaseGetUsers{
        {
            Cookie:       http.Cookie{},
            ResponsePart: `"error":"no cookie"`,
            StatusCode:   http.StatusUnauthorized,
        },
        {
            Cookie:       http.Cookie{Name: configs.CookieAuthName, Value: "fake"},
            ResponsePart: `"error":"no user with the cookie"`,
            StatusCode:   http.StatusUnauthorized,
        },
        {
            Cookie:       cookie,
            ResponsePart: `"user_id":1`,
            StatusCode:   http.StatusOK,
        },
    }

    for caseNum, item := range cases {
        url := "http://167.71.5.55/api/users/"
        req := httptest.NewRequest("GET", url, nil)
        req.AddCookie(&item.Cookie)
        w := httptest.NewRecorder()

        handler := GetHandlerUsers(db)
        handler(w, req)

        CheckStatusCodeAndResponse(t, caseNum, w, item.StatusCode, item.ResponsePart)
    }
}
