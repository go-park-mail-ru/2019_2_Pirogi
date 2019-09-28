package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/auth"
	Error "github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/error"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/images"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/inmemory"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

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
			_, _ = fmt.Fprint(w, Error.New(400, "invalid json", err.Error()))
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
		err := auth.Logout(w, r, db)
		if err != nil {
			Error.Render(w, Error.New(400, err.Error()))
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
		jsonBody, _ := json.Marshal(obj)
		_, err = fmt.Fprint(w, string(jsonBody), "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func GetHandlerUsersCreate(db *inmemory.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rawBody, err := ioutil.ReadAll(r.Body)

		if err != nil {
			Error.Render(w, Error.New(400, err.Error()))
			return
		}
		defer r.Body.Close()

		newUser := &models.NewUser{}
		err = newUser.UnmarshalJSON(rawBody)
		if err != nil {
			Error.Render(w, Error.New(400, err.Error()))
			return
		}

		e := db.Insert(newUser, 0)

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
			Error.Render(w, Error.New(http.StatusBadRequest, err.Error()))
			return
		}
		defer r.Body.Close()

		updateUser := &models.UpdateUser{}
		err = updateUser.UnmarshalJSON(rawBody)
		if err != nil {
			Error.Render(w, Error.New(http.StatusBadRequest, err.Error()))
			return
		}

		session, err := r.Cookie(configs.CookieAuthName)
		if err != nil {
			Error.Render(w, Error.New(401, err.Error()))
			return
		}
		userForCookie, ok := db.FindUserByCookie(*session)
		if !ok {
			Error.Render(w, Error.New(401, "no user with the cookie"))
			return
		}
		updateUser.ActualEmail = userForCookie.Email
		user, ok := db.FindByEmail(updateUser.ActualEmail)
		if !ok {
			Error.Render(w, Error.New(http.StatusNotFound, "no user with the email"))
			return
		}

		// я тупой и не смог сделать switch правильно. Признаю
		if updateUser.Name != "" {
			user.Name = updateUser.Name
		}
		if updateUser.Password != "" {
			user.Password = updateUser.Password
		}
		if updateUser.Email != "" {
			user.Email = updateUser.Email
			db.Insert(session, user.ID)
		}
		if updateUser.Description != "" {
			user.Description = updateUser.Description
		}

		e := db.Insert(user, 0)
		if e != nil {
			Error.Render(w, e)
			return
		}

	}
}

func UploadImageHandler(w http.ResponseWriter, r *http.Request) {
	ID, loadTarget, e := images.GetFields(r)
	if e != nil {
		Error.Render(w, e)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, images.MaxUploadSize)
	if err := r.ParseMultipartForm(images.MaxUploadSize); err != nil {
		Error.Render(w, Error.New(http.StatusBadRequest, err.Error()))
		return
	}
	defer r.Body.Close()

	file, _, err := r.FormFile("uploadFile")
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

	var uploadPath string
	switch loadTarget {
	case "user":
		uploadPath = configs.ImageUploadPath
	default:
		Error.Render(w, e)

	}
	fileName := images.GenerateFilename(loadTarget, ID, ending)
	e = images.WriteFile(fileBytes, fileName, uploadPath)
	if e != nil {
		Error.Render(w, e)
		return
	}
}
