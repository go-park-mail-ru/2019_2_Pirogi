package handlers

import (
	"encoding/json"
	"fmt"
	Error "github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/error"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/images"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/auth"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/inmemory"
	"github.com/gorilla/mux"
)

func GetHandlerLogin(db *inmemory.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rawBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			Error.Render(w, http.StatusBadRequest, err.Error())
			return
		}
		credentials := models.Credentials{}
		err = credentials.UnmarshalJSON(rawBody)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = fmt.Fprint(w, Error.Wrap("invalid json", err))
			return
		}
		err = auth.Login(w, r, db, credentials.Email, credentials.Password)
		if err != nil {
			Error.Render(w, http.StatusBadRequest, err.Error())
			return
		}
	}
}

func GetHandlerLogout(db *inmemory.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.Logout(w, r, db)
		if err != nil {
			Error.Render(w, http.StatusBadRequest, err.Error())
		}
	}
}

func GetHandlerUser(db *inmemory.DB) http.HandlerFunc {
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

func GetHandlerUsersCreate(db *inmemory.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rawBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			Error.Render(w, http.StatusBadRequest, err.Error())
			return
		}
		newUser := &models.NewUser{}
		err = newUser.UnmarshalJSON(rawBody)
		if err != nil {
			Error.Render(w, http.StatusBadRequest, err.Error())
			return
		}

		err = db.Insert(newUser)

		if err != nil {
			Error.Render(w, http.StatusBadRequest, "not able to insert into db: ", err.Error())
			return
		}
		err = auth.Login(w, r, db, newUser.Email, newUser.Password)
		if err != nil {
			Error.Render(w, http.StatusBadRequest, "not able to auth: ", err.Error())
			return
		}
	}
}

func GetHandlerUsersUpdate(db *inmemory.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func UploadImageHandler(w http.ResponseWriter, r *http.Request) {
	ID, loadTarget, err := images.GetFields(r)
	if err != nil {
		Error.Render(w, http.StatusBadRequest, err.Error())
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, images.MaxUploadSize)
	if err := r.ParseMultipartForm(images.MaxUploadSize); err != nil {
		Error.Render(w, http.StatusBadRequest, err.Error())
		return
	}

	file, _, err := r.FormFile("uploadFile")
	if err != nil {
		Error.Render(w, http.StatusBadRequest, err.Error())
		return
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		Error.Render(w, http.StatusBadRequest, err.Error())
		return
	}

	ending, err := images.DetectContentType(fileBytes)
	if err != nil {
		Error.Render(w, http.StatusBadRequest, err.Error())
		return
	}

	var uploadPath string
	switch loadTarget {
	case "user":
		uploadPath = images.UploadUsersPath
	default:
		Error.Render(w, http.StatusBadRequest, "set the target")

	}
	fileName := images.GenerateFilename(loadTarget, ID, ending)
	err = images.WriteFile(fileBytes, fileName, uploadPath)
	if err != nil {
		Error.Render(w, http.StatusBadRequest, err.Error())
		return
	}
}
