package inmemory

import (
	Error "github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/error"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/user"
	"net/http"
	"strconv"
)

type DB struct {
	users            map[int]models.User
	films            map[int]models.Film
	usersAuthCookies map[int]http.Cookie
}

func Init() *DB {
	users := make(map[int]models.User, 0)
	films := make(map[int]models.Film, 0)
	usersAuthCookies := make(map[int]http.Cookie, 0)
	db := DB{users: users, usersAuthCookies: usersAuthCookies, films: films}
	return &db
}

func (db *DB) GetID(target string) int {
	switch target {
	case "user":
		return len(db.users)
	case "film":
		return len(db.films)
	case "auth_cookie":
		return len(db.usersAuthCookies)
	default:
		return 0
	}
}

// затирает старые записи
func (db *DB) Insert(in interface{}) *models.Error {
	switch in.(type) {
	case models.NewUser:
		newUser := in.(models.NewUser)
		_, ok := db.FindByEmail(newUser.Email)
		if ok {
			return Error.New(400, "user with the email is already existed")
		}
		u, e := user.CreateNewUser(db.GetID("user"), newUser)
		if e != nil {
			return e
		}
		db.users[db.GetID("user")] = u
		return nil
	case models.User:
		u := in.(models.User)
		if _, ok := db.users[u.ID]; ok {
			db.users[u.ID] = u
			return nil
		}
		return Error.New(404, "user not found")
	case http.Cookie:
		cookie := in.(http.Cookie)
		db.usersAuthCookies[db.GetID("auth_cookie")] = cookie
		return nil
	default:
		return Error.New(400, "not supported type")
	}
}

func (db *DB) Delete(in interface{}) {
	switch in.(type) {
	case http.Cookie:
		cookie := in.(http.Cookie)
		u, ok := db.FindUserByCookie(cookie)
		if !ok {
			return
		}
		if _, ok := db.usersAuthCookies[u.ID]; !ok {
			db.usersAuthCookies[u.ID] = cookie
		}
	}
}

func (db *DB) Get(id int, target string) (interface{}, *models.Error) {
	switch target {
	case "user":
		if u, ok := db.users[id]; ok {
			return u, nil
		}
		return nil, Error.New(404, "no user with id: "+strconv.Itoa(id))
	case "film":
		if f, ok := db.films[id]; ok {
			return f, nil
		}
		return nil, Error.New(404, "no film with the id: "+strconv.Itoa(id))
	}
	return nil, Error.New(404, "not supported type: "+target)
}

func (db *DB) FindByEmail(email string) (models.User, bool) {
	for k, u := range db.users {
		if u.Email == email {
			return db.users[k], true
		}
	}
	return models.User{}, false
}

func (db *DB) FindByID(id int) (models.User, bool) {
	for k, u := range db.users {
		if u.ID == id {
			return db.users[k], true
		}
	}
	return models.User{}, false
}

func (db *DB) CheckCookie(cookie http.Cookie) bool {
	for _, v := range db.usersAuthCookies {
		if v.Value == cookie.Value {
			return true
		}
	}
	return false
}

func (db *DB) FindUserByCookie(cookie http.Cookie) (models.User, bool) {
	for k, v := range db.usersAuthCookies {
		if v.Value == cookie.Value {
			u, ok := db.FindByID(k)
			if !ok {
				return models.User{}, false
			}
			return u, true
		}
	}
	return models.User{}, false
}
