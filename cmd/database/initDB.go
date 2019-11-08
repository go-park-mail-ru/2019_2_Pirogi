package main

import (
	"flag"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/common"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/database"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/user"
	"github.com/labstack/gommon/log"
)

func upsertNewPerson(conn *database.MongoConnection, name, birthday, birthplace string) {
	conn.Upsert(models.NewPerson{
		Name:       name,
		Roles:      []models.Role{"actor"},
		Birthday:   birthday,
		Birthplace: birthplace,
	})
}

func upsertNewFilm(conn *database.MongoConnection, title, description, year string, countries []string, genres []models.Genre, ids []models.ID) {
	conn.Upsert(models.NewFilm{
		Title:       title,
		Description: description,
		Year:        year,
		Countries:   countries,
		Genres:      genres,
		PersonsID:   ids,
		Trailer:     "https://www.youtube.com/watch?v=oqHJp_ZZdU4",
	})
}

func FakeFillDB(conn *database.MongoConnection) {
	cookie := http.Cookie{
		Name:  configs.Default.CookieAuthName,
		Value: "value",
		Path:  "/",
	}

	conn.Upsert(models.NewUser{
		Credentials: models.Credentials{Email: "oleg@mail.ru", Password: user.GetMD5Hash("qwerty123")},
		Username:    "Oleg",
	})
	conn.Upsert(models.UserCookie{UserID: 0, Cookie: &cookie})

	conn.Upsert(models.NewUser{
		Credentials: models.Credentials{Email: "anton@mail.ru", Password: user.GetMD5Hash("qwe523")},
		Username:    "Anton",
	})
	conn.Upsert(models.UserCookie{UserID: 1, Cookie: &cookie})

	conn.Upsert(models.NewUser{
		Credentials: models.Credentials{Email: "yura@gmail.com", Password: user.GetMD5Hash("12312312")},
		Username:    "Yura",
	})
	conn.Upsert(models.UserCookie{UserID: 2, Cookie: &cookie})

	upsertNewPerson(conn, "Роберт Редфорд", "США", "1936")
	upsertNewPerson(conn, "Том Мюррэй", "США", "1973")
	upsertNewPerson(conn, "Елена Драпеко", "СССР", "1948")
	upsertNewPerson(conn, "Эдди Альберт", "США", "1906")
	upsertNewPerson(conn, "Юэн Макгрегор", "США", "1976")
	upsertNewPerson(conn, "Мэри МакДоннелл", "США", "1966")
	upsertNewPerson(conn, "Хейден Кристенсен", "США", "1986")
	upsertNewPerson(conn, "Феллипе Хаагенсен", "США", "1956")
	upsertNewPerson(conn, "Шон Пенн", "США", "1986")
	upsertNewPerson(conn, "Вуди Харрельсон", "США", "1966")
	upsertNewPerson(conn, "Мак Суэйн", "США", "1975")
	upsertNewPerson(conn, "Ольга Остроумова", "США", "1937")
	upsertNewPerson(conn, "Виталий Соломин", "СССР", "1989")
	upsertNewPerson(conn, "Мишель Пфайффер", "США", "1950")
	upsertNewPerson(conn, "Александр Борисов", "СССР", "1967")
	upsertNewPerson(conn, "Роберт Шоу", "США", "1977")
	upsertNewPerson(conn, "Роберт Де Ниро", "США", "1959")
	upsertNewPerson(conn, "Брюс Уиллис", "США", "1969")
	upsertNewPerson(conn, "Константин Сорокин", "СССР", "1949")
	upsertNewPerson(conn, "Тони Кертис", "США", "1990")
	upsertNewPerson(conn, "Хэйли Джоэл Осмент", "США", "1978")
	upsertNewPerson(conn, "Одри Хепберн", "США", "1979")
	upsertNewPerson(conn, "Леонид Быков", "СССР", "1935")
	upsertNewPerson(conn, "Джек Леммон", "США", "1976")
	upsertNewPerson(conn, "Дакота Фаннинг", "США", "1956")
	upsertNewPerson(conn, "Леандру Фирмину", "США", "1976")
	upsertNewPerson(conn, "Шарлиз Терон", "США", "1981")
	upsertNewPerson(conn, "Екатерина Маркова", "СССР", "1984")
	upsertNewPerson(conn, "Василий Ливанов", "США", "1956")
	upsertNewPerson(conn, "Киану Ривз", "США", "1976")
	upsertNewPerson(conn, "Кевин Костнен", "США", "1992")
	upsertNewPerson(conn, "Грэм Грин", "США", "1973")
	upsertNewPerson(conn, "Уилл Смит", "США", "1978")
	upsertNewPerson(conn, "Лоренс Фишбёрн", "США", "1989")
	upsertNewPerson(conn, "Кэрри-Энн Мосс", "США", "1982")
	upsertNewPerson(conn, "Мэрилин Монро", "США", "1975")
	upsertNewPerson(conn, "Грегори Пек", "США", "1987")
	upsertNewPerson(conn, "Кьюба Гудинг мл.", "США", "1986")
	upsertNewPerson(conn, "Чарльз Чаплин", "США", "1936")
	upsertNewPerson(conn, "Натали Портман", "США", "1946")
	upsertNewPerson(conn, "Розарио Доусон", "США", "1956")
	upsertNewPerson(conn, "Тони Коллетт", "США", "1976")
	upsertNewPerson(conn, "Никита Михалков", "СССР", "1976")

	conn.Upsert(models.NewReview{
		Title:    "Best review title",
		Body:     "Best review body",
		FilmID:   0,
		AuthorID: 1,
	})

	time.Sleep(2 * time.Second)

	conn.Upsert(models.NewReview{
		Title:    "Best review title 2",
		Body:     "Best review body 2",
		FilmID:   0,
		AuthorID: 2,
	})

	conn.Upsert(models.Stars{
		UserID: 0,
		FilmID: 0,
		Mark:   3.0,
	})

	conn.Upsert(models.Stars{
		UserID: 1,
		FilmID: 1,
		Mark:   5.0,
	})

	conn.Upsert(models.Like{
		UserID:   0,
		Target:   "film",
		TargetID: 0,
	})

	conn.Upsert(models.Like{
		UserID:   1,
		Target:   "review",
		TargetID: 0,
	})
}

func main() {
	configsPath := flag.String("config", "configs/", "directory with configs")
	flag.Parse()

	err := common.UnmarshalConfigs(configsPath)
	if err != nil {
		log.Fatal(err.Error())
	}

	conn, err := database.InitMongo("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}

	conn.ClearDB()
	err = conn.InitCounters()
	if err != nil {
		log.Fatal(err)
	}
	FakeFillDB(conn)
}
