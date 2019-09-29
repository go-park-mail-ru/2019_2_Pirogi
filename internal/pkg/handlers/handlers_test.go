package handlers

import (
    "github.com/go-park-mail-ru/2019_2_Pirogi/configs"
    "github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/inmemory"
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "testing"
)

func CheckStatusCodeAndResponse(t *testing.T, caseNumber int, w *httptest.ResponseRecorder, expectedCode int, expectedResponse string) {
    if w.Code != expectedCode {
        t.Errorf("[%d] wrong StatusCode: got %d, expected %d", caseNumber, w.Code, expectedCode)
    }

    resp := w.Result()
    body, _ := ioutil.ReadAll(resp.Body)

    bodyStr := string(body)
    if bodyStr != expectedResponse {
        t.Errorf("[%d] wrong Response: got %+v, expected %+v", caseNumber, bodyStr, expectedResponse)
    }
}

type TestCaseGetUser struct {
    ID         string
    Response   string
    StatusCode int
}

func TestGetUser(t *testing.T) {
    db := inmemory.Init()
    db.FakeFillDB()

    cases := []TestCaseGetUser{
        {
            ID:         "1",
            Response:   `{"status": 200, "resp": {"user": 1}}`,
            StatusCode: http.StatusOK,
        },
        {
            ID:         "500",
            Response:   `{"status": 500, "err": "db_error"}`,
            StatusCode: http.StatusInternalServerError,
        },
    }

    // It does not work yet
    for caseNum, item := range cases {
        url := "http://167.71.5.55/api/users/" + item.ID
        req := httptest.NewRequest("GET", url, nil)
        //ctx := context.WithValue(req.Context(), "user_id", item.ID)
        w := httptest.NewRecorder()

        handler := GetHandlerUser(db)
        handler(w, req)
        // handler(w, req.WithContext(ctx))

        CheckStatusCodeAndResponse(t, caseNum, w, item.StatusCode, item.Response)
    }
}

type TestCaseGetUsers struct {
    Cookie     http.Cookie
    Response   string
    StatusCode int
}

func TestGetUsers(t *testing.T) {
    db := inmemory.Init()
    db.FakeFillDB()
    cookie := http.Cookie{Name: configs.CookieAuthName, Value: "cookie"}
    db.Insert(cookie, 1)

    cases := []TestCaseGetUsers{
        {
            Cookie:     http.Cookie{},
            Response:   `{"status":401,"error":"no cookie"}`,
            StatusCode: http.StatusUnauthorized,
        },
        {
            Cookie:     http.Cookie{Name: configs.CookieAuthName, Value: "fake"},
            Response:   `{"status":401,"error":"no user with the cookie"}{"user_id":0,"name":"","rating":0,"description":"","avatar_link":"","email":"","password":""}`,
            StatusCode: http.StatusUnauthorized,
        },
        {
            Cookie:     cookie,
            Response:   `{"user_id":1,"name":"Anton","rating":8.3,"description":"","avatar_link":"anton.jpg","email":"anton@mail.ru","password":""}`,
            StatusCode: http.StatusOK,
        },
    }

    for caseNum, item := range cases {
        url := "http://167.71.5.55/api/users/"
        req := httptest.NewRequest("GET", url, nil)
        req.AddCookie(&item.Cookie)
        w := httptest.NewRecorder()

        handler := GetHandlerUsers(db)
        handler(w, req)

        CheckStatusCodeAndResponse(t, caseNum, w, item.StatusCode, item.Response)
    }
}
