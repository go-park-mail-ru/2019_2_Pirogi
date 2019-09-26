package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	Error "github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/error"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/auth"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/inmemory"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/user"
	"github.com/gorilla/mux"
)

func getObjectFromRequest(r *http.Request, t string) (interface{}, error) {
	rawBody, _ := ioutil.ReadAll(r.Body)
	defer func() { _ = r.Body.Close() }()
	var obj interface{}

	switch t {
	case "user":
		obj = new(user.User)
	case "credentials":
		obj = new(user.Credentials)
	}

	err := json.Unmarshal(rawBody, &obj)
	if err != nil {
		return nil, errors.New("error while parsing json: " + err.Error())
	}
	return obj, nil
}

func GetHandlerLogin(db *inmemory.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		jsonBody, err := getObjectFromRequest(r, "credentials")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = fmt.Fprint(w, Error.Wrap("invalid json", err))
			return
		}
		u := jsonBody.(*user.Credentials)
		err = auth.Auth(w, r, db, u.Email, u.Password)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = fmt.Fprint(w, Error.Wrap("invalid credentials", err))
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func GetHandlerLogout(db *inmemory.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.Logout(w, r, db)
		if err != nil {
			_, _ = fmt.Fprint(w, err.Error())
		}
	}
}

func GetHandlerImages(db *inmemory.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func GetHandlerUser(db *inmemory.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		id, err := strconv.Atoi(mux.Vars(r)["user_id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = fmt.Fprintf(w, Error.InvalidQueryArgument("user_id"))
			return
		}
		obj, err := db.Get(id, "user")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			_, _ = fmt.Fprint(w, Error.NotFound())
			return
		}
		jsonBody, _ := json.Marshal(obj)
		_, err = fmt.Fprint(w, string(jsonBody), "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func GetHandlerUsersCreate(db *inmemory.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")

		jsonBody, err := getObjectFromRequest(r, "user")

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = fmt.Fprint(w, Error.Wrap("invalid json", err))
			return
		}

		err = db.Insert(jsonBody)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = fmt.Fprint(w, Error.Wrap("can not create user", err))
		}

		err = auth.Auth(w, r, db, jsonBody.(*user.User).Email, jsonBody.(*user.User).Password)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = fmt.Fprint(w, Error.Wrap("can not auth", err))
		}
		return
	}
}

func HandleDefault(w http.ResponseWriter, r *http.Request) {
	log.Print(r.URL)
}
