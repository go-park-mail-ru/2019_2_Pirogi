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

	conn.InsertOrUpdate(models.NewFilm{
		Title: "Бойцовский клуб",
		Description: "Терзаемый хронической бессонницей и отчаянно пытающийся вырваться из мучительно скучной жизни " +
			"клерк встречает некоего Тайлера Дардена, харизматического торговца мылом с извращенной философией. Тайлер " +
			"уверен, что самосовершенствование — удел слабых, а саморазрушение — единственное, ради чего стоит жить.",
		Date:      "1999",
		Countries: []string{"США", "Германия"},
		Genres:    []models.Genre{"триллер", "драма", "криминал"},
		Actors: []models.PersonTrunc{{0, "Эдвард Нортон"}, {1, "Брэд Питт"},
			{2, "Хелена Бонем Картер"}},
		Directors:     []models.PersonTrunc{{3, "Дэвид Финчер"}},
		Producers:     []models.PersonTrunc{{4, "Росс Грэйсон Белл"}, {5, "Сиан Чаффин"}},
		Compositors:   []models.PersonTrunc{{6, "Даст Бразерс"}, {7, "Джон Кинг"}},
		Screenwriters: []models.PersonTrunc{{8, "Джим Улс"}, {9, "Чак Паланик"}},
		Poster:        models.Image{0, "matrix.png"},
		Images:        []models.Image{},
	})
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
