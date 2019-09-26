package inmemory

import (
	"../user"
	"errors"
	"strconv"
)

type DB struct {
	usersNumber int
	users       map[int]user.User
}

func Init() *DB {
	users := make(map[int]user.User, 0)
	db := DB{users: users}
	return &db
}

func (db *DB) Insert(in interface{}) error {
	if u, ok := in.(*user.User); ok {
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
