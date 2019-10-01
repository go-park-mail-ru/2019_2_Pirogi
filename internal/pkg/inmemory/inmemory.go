package inmemory

import (
	Error "github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/error"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/film"
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
			return Error.New(400, "user with the email already exists")
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
	case models.NewFilm:
		newFilm := in.(models.NewFilm)
		// It is supposed that there cannot be films with the same title
		_, ok := db.FindFilmByTitle(newFilm.Title)
		if ok {
			return Error.New(400, "film with the title already exists")
		}
		f, e := film.CreateNewFilm(db.GetID("film"), newFilm)
		if e != nil {
			return e
		}
		db.films[db.GetID("film")] = f
		return nil
	case models.Film:
		f := in.(models.Film)
		if _, ok := db.users[f.ID]; ok {
			db.films[f.ID] = f
			return nil
		}
		return Error.New(404, "film not found")
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

func (db *DB) FindUserByID(id int) (models.User, bool) {
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
			u, ok := db.FindUserByID(k)
			if !ok {
				return models.User{}, false
			}
			return u, true
		}
	}
	return models.User{}, false
}

func (db *DB) FindFilmByTitle(title string) (models.Film, bool) {
	for k, f := range db.films {
		if f.Title == title {
			return db.films[k], true
		}
	}
	return models.Film{}, false
}

func (db *DB) FakeFillDB() {
	db.Insert(models.NewUser{
		Credentials: models.Credentials{Email: "oleg@mail.ru", Password: user.GetMD5Hash("qwerty123")},
		Username:    "Oleg",
	})

	db.Insert(models.NewUser{
		Credentials: models.Credentials{Email: "anton@mail.ru", Password: user.GetMD5Hash("qwe523")},
		Username:    "Anton",
	})

	db.Insert(models.NewUser{
		Credentials: models.Credentials{Email: "yura@gmail.com", Password: user.GetMD5Hash("12312312")},
		Username:    "Yura",
	})

	db.Insert(models.NewFilm{FilmInfo: models.FilmInfo{
		Title: "Бойцовский клуб",
		Description: "Терзаемый хронической бессонницей и отчаянно пытающийся вырваться из мучительно скучной жизни " +
			"клерк встречает некоего Тайлера Дардена, харизматического торговца мылом с извращенной философией. Тайлер " +
			"уверен, что самосовершенствование — удел слабых, а саморазрушение — единственное, ради чего стоит жить.",
		Date:       "1999",
		Actors:     []string{"Брэд Питт", "Эдвард Нортон"},
		Genres:     []string{"Драма", "Боевик"},
		Directors:  []string{"Дэвид Финчер"},
		Rating:     9.1,
		Image:      "club.jpg",
		ReviewsNum: models.ReviewsNum{Total: 100, Positive: 90, Negative: 10},
	}})

	db.Insert(models.NewFilm{FilmInfo: models.FilmInfo{
		Title: "Матрица",
		Description: "Мир Матрицы — это иллюзия, существующая только в бесконечном сне обреченного человечества. " +
			"Холодный мир будущего, в котором люди — всего лишь батарейки в компьютерных системах.",
		Date:       "1999",
		Actors:     []string{"Киану Ривз", "Кэрри-Энн Мосс"},
		Genres:     []string{"Фэнтези"},
		Directors:  []string{"Лана Вачовски", "Лилли Вачовски"},
		Rating:     8.9,
		Image:      "matrix.jpg",
		ReviewsNum: models.ReviewsNum{Total: 110, Positive: 90, Negative: 20},
	}})
}
