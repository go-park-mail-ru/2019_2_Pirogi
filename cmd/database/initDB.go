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

func upsertNewFilm(conn *database.MongoConnection, title, description, year string, countries []string, genres []models.Genre, ids []models.ID, trailer string) {
	conn.Upsert(models.NewFilm{
		Title:       title,
		Description: description,
		Year:        year,
		Countries:   countries,
		Genres:      genres,
		PersonsID:   ids,
		Trailer:     trailer,
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
	upsertNewFilm(conn, "Семь жизней",
		"Инженер Бен отправляется в необычное путешествие. "+
			"В ходе своей поездки он встречает семерых незнакомцев, "+
			"включая смертельно больную Эмили, которая называет себя девушкой "+
			"с подбитыми крыльями. Бен неожиданно влюбляется в нее, "+
			"что сильно усложняет его первоначальный план. "+
			"Сможет ли он разгадать послание судьбы?",
		"2008", []string{"США"}, []models.Genre{"драма",
			"мелодрама"}, []models.ID{34, 42, 10}, "https://www.youtube.com/watch?v=3sO8hJoKV5A")

	upsertNewFilm(conn, "Город Бога",
		"Фильм охватывает события, происходящие на протяжении тридцати лет в так называемом «Городе Бога» — трущобах в бразильском городе Рио-де-Жанейро. Главные герои фильма — парень по кличке «Ракета», балансирующий между честной жизнью и мелкими правонарушениями и его знакомый Дадинью, который с восьми лет начал карьеру гангстера.",
		"2002", []string{"Бразилия"}, []models.Genre{"драма",
			"криминал"}, []models.ID{6, 27, 8}, "https://www.youtube.com/watch?v=rYUVcZH1-8U")

	upsertNewFilm(conn, "Приключения Шерлока Холмса и доктора Ватсона: Собака Баскервилей",
		"Труп Чарльза Баскервиля обнаруживают неподалеку от его родового поместья. Выражение нечеловеческого ужаса на лице покойника и следы крупной собаки поблизости заставляют вспомнить старинную легенду о проклятии, тяготеющем над родом Баскервилей. Шерлоку Холмсу предстоит докопаться до истины.",
		"1981", []string{"СССР"}, []models.Genre{"криминал",
			"детектив",
			"приключения"}, []models.ID{30, 13, 44}, "https://www.youtube.com/watch?v=zX0y-9UushI")
	upsertNewFilm(conn, "Матрица",
		"Жизнь Томаса Андерсона разделена на две части: днём он — самый обычный офисный работник, получающий нагоняи от начальства, а ночью превращается в хакера по имени Нео, и нет места в сети, куда он не смог бы дотянуться. Но однажды всё меняется — герой, сам того не желая, узнаёт страшную правду: всё, что его окружает — не более, чем иллюзия, Матрица, а люди — всего лишь источник питания для искусственного интеллекта, поработившего человечество. И только Нео под силу изменить расстановку сил в этом чужом и страшном мире.",
		"1999", []string{"США"}, []models.Genre{"фантастика",
			"боевик"}, []models.ID{31, 35, 36}, "https://www.youtube.com/watch?v=YihPA42fdQ8")

	upsertNewFilm(conn, "Военный ныряльщик",
		"Фильм основан на реальных событиях из жизни легендарного водолаза Карла Брашира. Его наставник и старший офицер Билл Сандэй убежден, что неграм нечего делать во флоте, и самыми жестокими и бесчеловечными способами издевается и обламывает амбициозного новичка.Однако ему это не удается, и более того, вскоре он начинает испытывать симпатию к этому чрезвычайно упертому парню, который скорее погибнет, чем покажет слабость. Вместе они пытаются сопротивляться бюрократам из командования ВМФ, которых не устраивает цвет кожи героя.Карл всегда вызывается добровольцем на выполнение самых опасных секретных спецопераций, связанных с повышенным риском для жизни. Однако судьба приготовила ему куда более страшное испытание, преодолеть которое способен только настоящий герой.",
		"2000", []string{"США"}, []models.Genre{"драма",
			"биография"}, []models.ID{17, 39, 28}, "https://www.youtube.com/watch?v=WAT_hOvMuoM")
	upsertNewFilm(conn, "Римские каникулы",
		"Инженер Бен отправляется в необычное путешествие. "+
			"В ходе своей поездки он встречает семерых незнакомцев, "+
			"включая смертельно больную Эмили, которая называет себя девушкой "+
			"с подбитыми крыльями. Бен неожиданно влюбляется в нее, "+
			"что сильно усложняет его первоначальный план. "+
			"Сможет ли он разгадать послание судьбы?",
		"1925", []string{"США"}, []models.Genre{"мелодрама",
			"комедия"}, []models.ID{38, 3}, "https://www.youtube.com/watch?v=xC_IuwO0R8c")
	upsertNewFilm(conn, "Золотая лихорадка",
		"Инженер Бен отправляется в необычное путешествие. "+
			"В ходе своей поездки он встречает семерых незнакомцев, "+
			"включая смертельно больную Эмили, которая называет себя девушкой "+
			"с подбитыми крыльями. Бен неожиданно влюбляется в нее, "+
			"что сильно усложняет его первоначальный план. "+
			"Сможет ли он разгадать послание судьбы?",
		"2008", []string{"США"}, []models.Genre{"драма",
			"комедия",
			"приключения",
			"семейный"}, []models.ID{40, 11}, "https://www.youtube.com/watch?v=5-0Qxoz6-Ls")
	upsertNewFilm(conn, "Я – Сэм",
		"Инженер Бен отправляется в необычное путешествие. "+
			"В ходе своей поездки он встречает семерых незнакомцев, "+
			"включая смертельно больную Эмили, которая называет себя девушкой "+
			"с подбитыми крыльями. Бен неожиданно влюбляется в нее, "+
			"что сильно усложняет его первоначальный план. "+
			"Сможет ли он разгадать послание судьбы?",
		"2001", []string{"США"}, []models.Genre{"драма"}, []models.ID{9, 14}, "")
	upsertNewFilm(conn, "В джазе только девушки",
		"Инженер Бен отправляется в необычное путешествие. "+
			"В ходе своей поездки он встречает семерых незнакомцев, "+
			"включая смертельно больную Эмили, которая называет себя девушкой "+
			"с подбитыми крыльями. Бен неожиданно влюбляется в нее, "+
			"что сильно усложняет его первоначальный план. "+
			"Сможет ли он разгадать послание судьбы?",
		"1959", []string{"США"}, []models.Genre{"мелодрама",
			"комедия",
			"криминал",
			"приключения",
			"музыка"}, []models.ID{37, 21}, "https://www.youtube.com/watch?v=a1KACVCmn2U")
	upsertNewFilm(conn, "...А зори здесь тихие",
		"Инженер Бен отправляется в необычное путешествие. "+
			"В ходе своей поездки он встречает семерых незнакомцев, "+
			"включая смертельно больную Эмили, которая называет себя девушкой "+
			"с подбитыми крыльями. Бен неожиданно влюбляется в нее, "+
			"что сильно усложняет его первоначальный план. "+
			"Сможет ли он разгадать послание судьбы?",
		"1972", []string{"CCCР"}, []models.Genre{"драма",
			"военный",
			"история"}, []models.ID{2, 29}, "https://www.youtube.com/watch?v=k2jYcjHtwWU")
	upsertNewFilm(conn, "Звёздные войны: Эпизод 3 – Месть Ситхов",
		"Инженер Бен отправляется в необычное путешествие. "+
			"В ходе своей поездки он встречает семерых незнакомцев, "+
			"включая смертельно больную Эмили, которая называет себя девушкой "+
			"с подбитыми крыльями. Бен неожиданно влюбляется в нее, "+
			"что сильно усложняет его первоначальный план. "+
			"Сможет ли он разгадать послание судьбы?",
		"2005", []string{"США"}, []models.Genre{"фантастика",
			"фэнтези",
			"боевик",
			"приключения"}, []models.ID{4, 41}, "https://www.youtube.com/watch?v=Uz_SI5JC5-Q")
	upsertNewFilm(conn, "Шестое чувство",
		"Инженер Бен отправляется в необычное путешествие. "+
			"В ходе своей поездки он встречает семерых незнакомцев, "+
			"включая смертельно больную Эмили, которая называет себя девушкой "+
			"с подбитыми крыльями. Бен неожиданно влюбляется в нее, "+
			"что сильно усложняет его первоначальный план. "+
			"Сможет ли он разгадать послание судьбы?",
		"1999", []string{"США"}, []models.Genre{"триллер",
			"драма",
			"детектив"}, []models.ID{19, 43}, "https://www.youtube.com/watch?v=1k2kkLxtErU")
	upsertNewFilm(conn, "Афера",
		"Инженер Бен отправляется в необычное путешествие. "+
			"В ходе своей поездки он встречает семерых незнакомцев, "+
			"включая смертельно больную Эмили, которая называет себя девушкой "+
			"с подбитыми крыльями. Бен неожиданно влюбляется в нее, "+
			"что сильно усложняет его первоначальный план. "+
			"Сможет ли он разгадать послание судьбы?",
		"1973", []string{"США"}, []models.Genre{"драма",
			"комедия",
			"криминал"}, []models.ID{18, 0}, "https://www.youtube.com/watch?v=H8c9UDohboU")
	upsertNewFilm(conn, "Максим Перепелица",
		"Инженер Бен отправляется в необычное путешествие. "+
			"В ходе своей поездки он встречает семерых незнакомцев, "+
			"включая смертельно больную Эмили, которая называет себя девушкой "+
			"с подбитыми крыльями. Бен неожиданно влюбляется в нее, "+
			"что сильно усложняет его первоначальный план. "+
			"Сможет ли он разгадать послание судьбы?",
		"1955", []string{"США"}, []models.Genre{"комедия"}, []models.ID{24, 15}, "https://www.youtube.com/watch?v=-J0jSsn4x7c")
	upsertNewFilm(conn, "Танцующий с волками",
		"Инженер Бен отправляется в необычное путешествие. "+
			"В ходе своей поездки он встречает семерых незнакомцев, "+
			"включая смертельно больную Эмили, которая называет себя девушкой "+
			"с подбитыми крыльями. Бен неожиданно влюбляется в нее, "+
			"что сильно усложняет его первоначальный план. "+
			"Сможет ли он разгадать послание судьбы?",
		"1990", []string{"США"}, []models.Genre{"драма",
			"приключения",
			"вестерн",
			"военный",
			"история"}, []models.ID{32, 33}, "https://www.youtube.com/watch?v=kN6J52p9ksE")

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
