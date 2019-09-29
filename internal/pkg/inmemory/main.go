package inmemory

import (
	Error "github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/error"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/film"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/images"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/user"
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

// пока in-memory details нужны чтобы записать куку, потом уберу
func (db *DB) Insert(in interface{}, id int) *models.Error {
	switch in.(type) {

	// при регистрации
	case *models.NewUser:
		newUserDetails := in.(*models.NewUser)
		if _, ok := db.FindByEmail(newUserDetails.Email); !ok {
			newUser := &models.User{
				Credentials: newUserDetails.Credentials,
				ID:          len(db.users),
				Name:        newUserDetails.Name,
				Rating:      0,
				AvatarLink:  images.GenerateFilename("user", strconv.Itoa(len(db.users)), ".jpeg"),
			}
			db.users[len(db.users)] = *newUser
			return nil
		} else {
			return Error.New(400, "user is already exist")
		}
	//	при обновлении информации
	case models.User:
		updatedUser := in.(models.User)
		if _, ok := db.users[updatedUser.ID]; ok {
			db.users[updatedUser.ID] = updatedUser
			return nil
		}
		return Error.New(404, "user not found")
	case http.Cookie:
		cookie := in.(http.Cookie)
		db.usersAuthCookies[id] = cookie
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
	return nil, Error.New(404, "no such type: "+target)
}

func (db *DB) FakeFillDB() {
	newUser, _ := user.CreateUser(len(db.users), "oleg@mail.ru", "Oleg", user.GetMD5Hash("qwerty123"), "oleg.jpg", 7.3)
	db.users[len(db.users)] = newUser
	newUser, _ = user.CreateUser(len(db.users), "anton@mail.ru", "Anton", user.GetMD5Hash("qwe523"), "anton.jpg", 8.3)
	db.users[len(db.users)] = newUser
	newUser, _ = user.CreateUser(len(db.users), "yura@gmail.com", "Yura", user.GetMD5Hash("12312312"), "yura.jpg", 9.5)
	db.users[len(db.users)] = newUser

	fightClub, _ := film.CreateFilm("Бойцовский клуб", "Терзаемый хронической бессонницей и отчаянно"+
		"пытающийся вырваться из мучительно скучной жизни клерк встречает некоего Тайлера Дардена, харизматического "+
		"торговца мылом с извращенной философией. Тайлер уверен, что самосовершенствование — удел слабых, "+
		"а саморазрушение — единственное, ради чего стоит жить.", []string{"Драма", "Боевик"}, []string{"Брэд Питт", "Эдвард Нортон"},
		[]string{"Дэвид Финчер"})
	matrix, _ := film.CreateFilm("Матрица", "Мир Матрицы — это иллюзия, существующая только в"+
		" бесконечном сне обреченного человечества. Холодный мир будущего, в котором люди — всего лишь батарейки в"+
		" компьютерных системах.", []string{"Фэнтези"}, []string{"Киану Ривз", "Кэрри-Энн Мосс"},
		[]string{"Лана Вачовски", "Лилли Вачовски"})

	// не сделал insert для фильмов
	db.films[len(db.films)] = fightClub
	db.films[len(db.films)] = matrix

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
