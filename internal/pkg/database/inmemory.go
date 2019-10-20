package database

import (
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"

	Error "github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/error"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/film"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/user"
)

type InmemoryDB struct {
	users            map[int]models.User
	films            map[int]models.Film
	usersAuthCookies map[int]http.Cookie
}

func InitInmemory() *InmemoryDB {
	users := make(map[int]models.User)
	films := make(map[int]models.Film)
	usersAuthCookies := make(map[int]http.Cookie)
	db := InmemoryDB{users: users, usersAuthCookies: usersAuthCookies, films: films}
	return &db
}

func (db *InmemoryDB) GetIDForInsert(target string) int {
	switch target {
	case configs.UserTargetName:
		return len(db.users)
	case configs.FilmTargetName:
		return len(db.films)
	case configs.CookieTargetName:
		return len(db.usersAuthCookies)
	default:
		return 0
	}
}

// затирает старые записи
func (db *InmemoryDB) Insert(in interface{}) *models.Error {
	switch in := in.(type) {
	case models.NewUser:
		_, ok := db.FindUserByEmail(in.Email)
		if ok {
			return Error.New(400, "user with the email already exists")
		}
		u, e := user.CreateNewUser(db.GetIDForInsert(configs.UserTargetName), &in)
		if e != nil {
			return e
		}
		db.users[db.GetIDForInsert(configs.UserTargetName)] = u
	case models.User:
		if _, ok := db.users[in.ID]; ok {
			db.users[in.ID] = in
			return nil
		}
		return Error.New(404, "user not found")
	case models.NewFilm:
		// It is supposed that there cannot be films with the same title
		_, ok := db.FindFilmByTitle(in.Title)
		if ok {
			return Error.New(400, "film with the title already exists")
		}
		f, e := film.CreateNewFilm(db.GetIDForInsert(configs.FilmTargetName), &in)
		if e != nil {
			return e
		}
		db.films[db.GetIDForInsert(configs.FilmTargetName)] = f
	case models.Film:
		if _, ok := db.users[in.ID]; ok {
			db.films[in.ID] = in
			return nil
		}
		return Error.New(404, "film not found")
	case models.UserCookie:
		db.usersAuthCookies[in.UserID] = *in.Cookie
	default:
		return Error.New(400, "not supported type")
	}
	return nil
}

func (db *InmemoryDB) Get(id int, target string) (interface{}, *models.Error) {
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

func (db *InmemoryDB) CheckCookie(cookie *http.Cookie) bool {
	for i := range db.usersAuthCookies {
		if db.usersAuthCookies[i].Value == cookie.Value {
			return true
		}
	}
	return false
}

func (db *InmemoryDB) DeleteCookie(in interface{}) {
	switch in := in.(type) {
	case http.Cookie:
		u, ok := db.FindUserByCookie(&in)
		if !ok {
			return
		}
		delete(db.usersAuthCookies, u.ID)
	}
}

func (db *InmemoryDB) FindUserByEmail(email string) (models.User, bool) {
	for k, u := range db.users {
		if u.Email == email {
			return db.users[k], true
		}
	}
	return models.User{}, false
}

func (db *InmemoryDB) FindUserByID(id int) (models.User, bool) {
	for k, u := range db.users {
		if u.ID == id {
			return db.users[k], true
		}
	}
	return models.User{}, false
}

func (db *InmemoryDB) FindUserByCookie(cookie *http.Cookie) (models.User, bool) {
	for k, v := range db.usersAuthCookies {
		if v.Value == cookie.Value {
			return db.FindUserByID(k)
		}
	}
	return models.User{}, false
}

func (db *InmemoryDB) FindFilmByTitle(title string) (models.Film, bool) {
	for k := range db.films {
		if db.films[k].Title == title {
			return db.films[k], true
		}
	}
	return models.Film{}, false
}

func (db *InmemoryDB) FindFilmByID(id int) (models.Film, bool) {
	for k, f := range db.films {
		if f.ID == id {
			return db.films[k], true
		}
	}
	return models.Film{}, false
}

func (db *InmemoryDB) FakeFillDB() {
	cookie := http.Cookie{
		Name:  "cinsear_session",
		Value: "value",
		Path:  "/",
	}

	db.Insert(models.NewUser{
		Credentials: models.Credentials{Email: "oleg@mail.ru", Password: user.GetMD5Hash("qwerty123")},
		Username:    "Oleg",
	})
	db.Insert(models.UserCookie{UserID: 0, Cookie: &cookie})

	db.Insert(models.NewUser{
		Credentials: models.Credentials{Email: "anton@mail.ru", Password: user.GetMD5Hash("qwe523")},
		Username:    "Anton",
	})
	db.Insert(models.UserCookie{UserID: 1, Cookie: &cookie})

	db.Insert(models.NewUser{
		Credentials: models.Credentials{Email: "yura@gmail.com", Password: user.GetMD5Hash("12312312")},
		Username:    "Yura",
	})
	db.Insert(models.UserCookie{UserID: 2, Cookie: &cookie})

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
