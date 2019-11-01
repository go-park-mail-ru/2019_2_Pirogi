package main

import (
	"flag"
	"net/http"

	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/common"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/database"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/user"
	"github.com/labstack/gommon/log"
)

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

	conn.Upsert(models.NewFilm{
		Title: "Бойцовский клуб",
		Description: "Терзаемый хронической бессонницей и отчаянно пытающийся вырваться из мучительно скучной жизни " +
			"клерк встречает некоего Тайлера Дардена, харизматического торговца мылом с извращенной философией. Тайлер " +
			"уверен, что самосовершенствование — удел слабых, а саморазрушение — единственное, ради чего стоит жить.",
		Year:      "1999",
		Countries: []string{"США", "Германия"},
		Genres:    []models.Genre{"триллер", "драма", "криминал"},
		Actors: []models.PersonTrunc{{0, "Эдвард Нортон", 4.5}, {1, "Брэд Питт", 4.8},
			{2, "Хелена Бонем Картер", 4.0}},
		Directors:     []models.PersonTrunc{{3, "Дэвид Финчер", 3.8}},
		Producers:     []models.PersonTrunc{{4, "Росс Грэйсон Белл", 4.7}, {5, "Сиан Чаффин", 4.5}},
		Compositors:   []models.PersonTrunc{{6, "Даст Бразерс", 4.1}, {7, "Джон Кинг", 4.2}},
		Screenwriters: []models.PersonTrunc{{8, "Джим Улс", 3.3}, {9, "Чак Паланик", 4.7}},
		Poster:        models.Image{0, "matrix.png"},
		Images:        []models.Image{},
	})

	conn.Upsert(models.NewReview{
		Title:    "Best review title",
		Body:     "Best review body",
		FilmID:   0,
		AuthorID: 1,
	})

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
		FilmID: 0,
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
	configsPath := flag.String("config", "../../configs/", "directory with configs")
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
