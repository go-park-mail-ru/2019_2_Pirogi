package handlers

import (
	"../error"
	"../inmemory"
	"../user"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func getObjectFromRequest(r *http.Request, t string) (interface{}, error) {
	rawBody, _ := ioutil.ReadAll(r.Body)
	defer func() { _ = r.Body.Close() }()
	var obj interface{}

	switch t {
	case "user":
		obj = new(user.User)
	}

	err := json.Unmarshal(rawBody, &obj)
	if err != nil {
		return nil, errors.New("error while parsing json: " + err.Error())
	}
	return obj, nil
}

func GetHandlerUser(db *inmemory.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")

		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = fmt.Fprint(w, Error.InvalidMethod(r.Method))
			return
		}
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
		_, err = fmt.Fprint(w, obj, "\n")
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
		return
	}
}

func HandleDefault(w http.ResponseWriter, r *http.Request) {
	log.Print(r.URL)
}
