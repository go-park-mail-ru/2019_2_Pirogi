package inmemory

import (
	"errors"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/images"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/user"
)

type DB struct {
	usersNumber      int
	users            map[int]user.User
	usersAuthCookies map[int]http.Cookie
}

func Init() *DB {
	users := make(map[int]user.User, 0)
	db := DB{users: users}
	return &db
}

func (db *DB) Insert(in interface{}) error {
	if u, ok := in.(*user.User); ok {
		u.Rating = 0.0
		u.AvatarLink = images.GenerateFilename("user", strconv.Itoa(u.ID), ".jpeg")
		(*u).ID = db.usersNumber
		db.users[db.usersNumber] = *u
		db.usersNumber++
		return nil
	}
	return errors.New("can not insert this type of object")
}

func (db *DB) Get(id int, t string) (interface{}, error) {
	switch t {
	case "user":
		if u, ok := db.users[id]; ok {
			return u, nil
		}
		return nil, errors.New("no user with id: " + strconv.Itoa(id))
	}
	return nil, errors.New("no such type: " + t)
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

func (db *DB) FindByEmail(email string) (user.User, bool) {
	for k, u := range db.users {
		if u.Credentials.Email == email {
			return db.users[k], true
		}
	}
	return user.User{}, false
}

func (db *DB) CheckCookie(userID int, cookie http.Cookie) bool {
	if c, ok := db.usersAuthCookies[userID]; ok {
		if c.Value == cookie.Value {
			return true
		}
	}
	return false
}

func (db *DB) DeleteCookie(cookie *http.Cookie) {
	for k, v := range db.usersAuthCookies {
		if v.Value == cookie.Value {
			delete(db.usersAuthCookies, k)
			break
		}
	}
}
