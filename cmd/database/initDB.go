package main

import (
	"flag"
	"net/http"
	"strconv"
	"time"

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

	conn.Upsert(models.NewPerson{
		Name:       "Эдвард Нортон",
		Roles:      []models.Role{"actor"},
		Birthday:   "18.08.1969",
		Birthplace: "Бостон",
	})

	conn.Upsert(models.NewPerson{
		Name:       "Брэд Питт",
		Roles:      []models.Role{"actor"},
		Birthday:   "18.10.1969",
		Birthplace: "Екатеринбург",
	})

	conn.Upsert(models.NewPerson{
		Name:       "Хелена Бонем Картер",
		Roles:      []models.Role{"actor"},
		Birthday:   "18.08.1989",
		Birthplace: "Лондон",
	})

	conn.Upsert(models.NewFilm{
		Title: "Бойцовский клуб",
		Description: "Терзаемый хронической бессонницей и отчаянно пытающийся вырваться из мучительно скучной жизни " +
			"клерк встречает некоего Тайлера Дардена, харизматического торговца мылом с извращенной философией. Тайлер " +
			"уверен, что самосовершенствование — удел слабых, а саморазрушение — единственное, ради чего стоит жить.",
		Year:      "1999",
		Countries: []string{"США", "Германия"},
		Genres:    []models.Genre{"триллер", "драма", "криминал"},
		PersonsID: []models.ID{0, 1, 2},
	})

	conn.Upsert(models.NewFilm{
		Title: "Матрица",
		Description: "Терзаемый хронической бессонницей и отчаянно пытающийся вырваться из мучительно скучной жизни " +
			"клерк встречает некоего Тайлера Дардена, харизматического торговца мылом с извращенной философией. Тайлер " +
			"уверен, что самосовершенствование — удел слабых, а саморазрушение — единственное, ради чего стоит жить.",
		Year:      "2009",
		Countries: []string{"Беларусь", "Германия"},
		Genres:    []models.Genre{"триллер", "драма", "криминал"},
		PersonsID: []models.ID{0, 1, 2},
	})

	for i := 3; i < 11; i++ {
		conn.Upsert(models.NewFilm{
			Title:       "Film " + strconv.Itoa(i),
			Description: "Description of film " + strconv.Itoa(i),
			Year:        "2009",
			Countries:   []string{"Беларусь", "Германия"},
			Genres:      []models.Genre{"триллер", "драма"},
			PersonsID:   []models.ID{0, 1, 2},
		})
	}

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
	configsPath := flag.String("config", "./configs/", "directory with configs")
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
