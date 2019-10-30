package main

import (
	"flag"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/common"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/database"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/user"
	"github.com/labstack/gommon/log"
	"net/http"
)

func FakeFillDB(conn *database.MongoConnection) {
	cookie := http.Cookie{
		Name:  configs.Default.CookieAuthName,
		Value: "value",
		Path:  "/",
	}

	conn.InsertOrUpdate(models.NewUser{
		Credentials: models.Credentials{Email: "oleg@mail.ru", Password: user.GetMD5Hash("qwerty123")},
		Username:    "Oleg",
	})
	conn.InsertOrUpdate(models.UserCookie{UserID: 0, Cookie: &cookie})

	conn.InsertOrUpdate(models.NewUser{
		Credentials: models.Credentials{Email: "anton@mail.ru", Password: user.GetMD5Hash("qwe523")},
		Username:    "Anton",
	})
	conn.InsertOrUpdate(models.UserCookie{UserID: 1, Cookie: &cookie})

	conn.InsertOrUpdate(models.NewUser{
		Credentials: models.Credentials{Email: "yura@gmail.com", Password: user.GetMD5Hash("12312312")},
		Username:    "Yura",
	})
	conn.InsertOrUpdate(models.UserCookie{UserID: 2, Cookie: &cookie})

	conn.InsertOrUpdate(models.NewFilm{FilmInfo: models.FilmInfo{
		Title: "Бойцовский клуб",
		Description: "Терзаемый хронической бессонницей и отчаянно пытающийся вырваться из мучительно скучной жизни " +
			"клерк встречает некоего Тайлера Дардена, харизматического торговца мылом с извращенной философией. Тайлер " +
			"уверен, что самосовершенствование — удел слабых, а саморазрушение — единственное, ради чего стоит жить.",
		Date:        "1999",
		ActorsID:    []models.ID{0, 1, 2, 3},
		GenresID:    []models.ID{2, 3},
		DirectorsID: []models.ID{0, 3},
		Rating:      9.1,
		ImagesID:    []models.ID{0, 3, 4, 5},
		ReviewsNum:  models.ReviewsNum{Total: 100, Positive: 90, Negative: 10},
	}})
	conn.InsertOrUpdate(models.NewFilm{FilmInfo: models.FilmInfo{
		Title: "Матрица",
		Description: "Мир Матрицы — это иллюзия, существующая только в бесконечном сне обреченного человечества. " +
			"Холодный мир будущего, в котором люди — всего лишь батарейки в компьютерных системах.",
		Date:        "1999",
		ActorsID:    []models.ID{0, 1, 2, 3},
		GenresID:    []models.ID{2, 3},
		DirectorsID: []models.ID{0, 3},
		Rating:      9.1,
		ImagesID:    []models.ID{0, 3, 4, 5},
		ReviewsNum:  models.ReviewsNum{Total: 100, Positive: 90, Negative: 10},
	}})
}

func main() {
	configsPath := flag.String("config", "../../configs/", "directory with configs")
	flag.Parse()

	err := common.UnmarshalConfigs(configsPath)
	if err != nil {
		log.Fatal(err.Error())
	}

	conn, err := database.InitMongo()
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
