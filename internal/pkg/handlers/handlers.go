package handlers

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/auth"
	Error "github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/error"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/images"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/inmemory"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/gorilla/mux"
)

func GetHandlerFilm(db *inmemory.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["film_id"])
		if err != nil {
			Error.Render(w, Error.New(404, "no film with the id", err.Error()))
			return
		}
		obj, e := db.Get(id, "film")
		if e != nil {
			Error.Render(w, e)
			return
		}
		film := obj.(models.Film)
		jsonBody, err := film.MarshalJSON()
		if err != nil {
			Error.Render(w, Error.New(500, "error while marshalling json", err.Error()))
			return
		}
		_, err = w.Write(jsonBody)
		if err != nil {
			Error.Render(w, Error.New(500, "error while writing response", err.Error()))
			return
		}
	}
}

func GetHandlerLoginCheck(db *inmemory.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ok := auth.LoginCheck(w, r, db)
		if !ok {
			w.WriteHeader(401)
			return
		}
	}
}

func GetHandlerLogin(db *inmemory.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rawBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			Error.Render(w, Error.New(500, err.Error()))
			return
		}
		defer r.Body.Close()
		credentials := models.Credentials{}
		err = credentials.UnmarshalJSON(rawBody)
		if err != nil {
			Error.Render(w, Error.New(400, "invalid json", err.Error()))
			return
		}
		e := auth.Login(w, r, db, credentials.Email, credentials.Password)
		if e != nil {
			Error.Render(w, e)
			return
		}
	}
}

func GetHandlerLogout(db *inmemory.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		e := auth.Logout(w, r, db)
		if e != nil {
			Error.Render(w, e)
		}
	}
}

func GetHandlerUser(db *inmemory.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["user_id"])
		if err != nil {
			Error.Render(w, Error.New(400, "invalid id", err.Error()))
			return
		}
		obj, e := db.Get(id, "user")
		if e != nil {
			Error.Render(w, e)
			return
		}
		user := obj.(models.User)
		userInfo := user.UserInfo
		jsonBody, err := userInfo.MarshalJSON()
		if err != nil {
			Error.Render(w, Error.New(500, err.Error()))
		}
		_, err = w.Write(jsonBody)
		if err != nil {
			Error.Render(w, Error.New(500, err.Error()))
		}
	}
}

func GetHandlerUsers(db *inmemory.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := r.Cookie(configs.CookieAuthName)
		if err != nil {
			Error.Render(w, Error.New(401, "no cookie"))
			return
		}
		user, ok := db.FindUserByCookie(*session)
		if !ok {
			Error.Render(w, Error.New(401, "invalid cookie"))
			return
		}
		rawUser, err := user.MarshalJSON()
		if err != nil {
			Error.Render(w, Error.New(500, err.Error()))
			return
		}
		_, err = w.Write(rawUser)
		if err != nil {
			Error.Render(w, Error.New(500, err.Error()))
			return
		}
	}
}

func GetHandlerUsersCreate(db *inmemory.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(configs.CookieAuthName)
		if err == nil && cookie != nil {
			if _, ok := db.FindUserByCookie(*cookie); ok {
				Error.Render(w, Error.New(403, "user is already logged in"))
				return
			}
		}
		rawBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			Error.Render(w, Error.New(400, err.Error()))
			return
		}
		defer r.Body.Close()
		newUser := models.NewUser{}
		err = newUser.UnmarshalJSON(rawBody)
		if err != nil {
			Error.Render(w, Error.New(400, err.Error()))
			return
		}
		e := db.Insert(newUser)
		if e != nil {
			Error.Render(w, e)
			return
		}
		e = auth.Login(w, r, db, newUser.Email, newUser.Password)
		if e != nil {
			Error.Render(w, e)
			return
		}
	}
}

func GetHandlerUsersUpdate(db *inmemory.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rawBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			Error.Render(w, Error.New(400, err.Error()))
			return
		}
		defer r.Body.Close()
		updateUser := &models.UserInfo{}
		err = updateUser.UnmarshalJSON(rawBody)
		if err != nil {
			Error.Render(w, Error.New(400, err.Error()))
			return
		}
		session, err := r.Cookie(configs.CookieAuthName)
		if err != nil {
			Error.Render(w, Error.New(401, err.Error()))
			return
		}
		user, ok := db.FindUserByCookie(*session)
		if !ok {
			Error.Render(w, Error.New(401, "no user with the cookie"))
			return
		}
		switch {
		case updateUser.Username != "":
			user.Username = updateUser.Username
			fallthrough
		case updateUser.Description != "":
			user.Description = updateUser.Description
		}
		e := db.Insert(user)
		if e != nil {
			Error.Render(w, e)
			return
		}
	}
}

func GetUploadImageHandler(db *inmemory.DB, target string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := r.Cookie(configs.CookieAuthName)
		if err != nil {
			Error.Render(w, Error.New(401, err.Error()))
			return
		}
		user, ok := db.FindUserByCookie(*session)
		if !ok {
			Error.Render(w, Error.New(401, "invalid cookie"))
			return
		}

		r.Body = http.MaxBytesReader(w, r.Body, images.MaxUploadSize)
		if err := r.ParseMultipartForm(images.MaxUploadSize); err != nil {
			Error.Render(w, Error.New(http.StatusBadRequest, err.Error()))
			return
		}
		defer r.Body.Close()

		file, _, err := r.FormFile("file")
		if err != nil {
			Error.Render(w, Error.New(http.StatusBadRequest, err.Error()))
			return
		}
		defer file.Close()

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			Error.Render(w, Error.New(http.StatusBadRequest, err.Error()))
			return
		}

		ending, e := images.DetectContentType(fileBytes)
		if e != nil {
			Error.Render(w, e)
			return
		}

		var path string
		fileName := images.GenerateFilename(time.Now().String(), strconv.Itoa(user.ID), ending)
		if target == "films" {
			path = configs.UsersImageUploadPath
		} else {
			path = configs.FilmsImageUploadPath
		}
		e = images.WriteFile(fileBytes, fileName, path)
		if e != nil {
			Error.Render(w, e)
			return
		}
		if target != "users" {
			return
		}
		user.Image = fileName
		e = db.Insert(user)
		if e != nil {
			Error.Render(w, e)
			return
		}
	}
}
