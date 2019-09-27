package inmemory

import (
	"errors"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/images"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/user"
)

type DB struct {
	usersNumber      int
	users            map[int]models.User
	usersAuthCookies map[string]http.Cookie
}

func Init() *DB {
	users := make(map[int]models.User, 0)
	usersAuthCookies := make(map[string]http.Cookie, 0)
	db := DB{users: users, usersAuthCookies: usersAuthCookies}
	return &db
}

func (db *DB) Insert(in interface{}) error {
	switch in.(type) {
	case *models.NewUser:
		newUserDetails := in.(*models.NewUser)
		if _, ok := db.FindByEmail(newUserDetails.Email); !ok {
			newUser := &models.User{
				Credentials: newUserDetails.Credentials,
				ID:          db.usersNumber,
				Name:        newUserDetails.Name,
				Rating:      0,
				AvatarLink:  images.GenerateFilename("user", strconv.Itoa(db.usersNumber), ".jpeg"),
			}
			db.users[db.usersNumber] = *newUser
			db.usersNumber++
			return nil
		}
		return errors.New("user is already existed")
	case http.Cookie:
		cookie := in.(http.Cookie)
		if _, ok := db.usersAuthCookies[cookie.Value]; !ok {
			db.usersAuthCookies[cookie.Value] = cookie
			return nil
		}
	default:
		return errors.New("not supported type")
	}
	return nil
}

func (db *DB) Delete(in interface{}) {
	switch in.(type) {
	case http.Cookie:
		cookie := in.(http.Cookie)
		if _, ok := db.usersAuthCookies[cookie.Value]; !ok {
			db.usersAuthCookies[cookie.Value] = cookie
		}
	}
}

func (db *DB) Get(id int, target string) (interface{}, error) {
	switch target {
	case "user":
		if u, ok := db.users[id]; ok {
			return u, nil
		}
		return nil, errors.New("no user with id: " + strconv.Itoa(id))
	}
	return nil, errors.New("no such type: " + target)
}

func (db *DB) FakeFillDB() {
	oleg, _ := user.CreateUser(0, "oleg@mail.ru", "Oleg", user.GetMD5Hash("qwerty123"), "oleg.jpg", 7.3)
	anton, _ := user.CreateUser(1, "anton@mail.ru", "Anton", user.GetMD5Hash("qwe523"), "anton.jpg", 8.3)
	yura, _ := user.CreateUser(2, "yura@gmail.com", "Yura", user.GetMD5Hash("12312312"), "yura.jpg", 9.5)
	db.users[oleg.ID] = oleg
	db.users[anton.ID] = anton
	db.users[yura.ID] = yura
	db.usersNumber = 3
}

func (db *DB) FindByEmail(email string) (models.User, bool) {
	for k, u := range db.users {
		if u.Credentials.Email == email {
			return db.users[k], true
		}
	}
	return models.User{}, false
}

func (db *DB) CheckCookie(email string, cookie http.Cookie) bool {
	if c, ok := db.usersAuthCookies[email]; ok {
		if c.Value == cookie.Value {
			return true
		}
	}
	return false
}
