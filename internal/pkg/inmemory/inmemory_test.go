package inmemory

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/user"
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFakeFillDB(t *testing.T) {
	db := Init()
	db.FakeFillDB()
	require.Equal(t, 3, len(db.users))
}

func TestDB_Insert(t *testing.T) {
	db := Init()
	ksyusha := models.NewUser{
		Credentials: models.Credentials{
			Email:    "ksyushag@mail.ru",
			Password: user.GetMD5Hash("qwerty123"),
		},
		Username: "Ksyusha",
	}
	err := db.Insert(ksyusha)
	require.Nil(t, err)
	require.True(t, reflect.DeepEqual(db.users[0].Credentials, ksyusha.Credentials))
}

func TestDB_InsertCookie(t *testing.T) {
	db := Init()
	cookie := http.Cookie{
		Name:  "cinsear_session",
		Value: "value",
		Path:  "/",
	}
	err := db.InsertCookie(cookie, 0)
	require.Nil(t, err)
	require.True(t, reflect.DeepEqual(cookie, db.usersAuthCookies[0]))
}

func TestDB_FindUserByCookie(t *testing.T) {
	db := Init()
	cookie := http.Cookie{
		Name:  "cinsear_session",
		Value: "value",
		Path:  "/",
	}
	err := db.InsertCookie(cookie, 0)
	require.Nil(t, err)
	ksyusha := models.NewUser{
		Credentials: models.Credentials{
			Email:    "ksyushag@mail.ru",
			Password: user.GetMD5Hash("qwerty123"),
		},
		Username: "Ksyusha",
	}
	err = db.Insert(ksyusha)
	require.Nil(t, err)
	foundUser, ok := db.FindUserByCookie(cookie)
	require.True(t, ok)
	require.Equal(t, ksyusha.Email, foundUser.Email)
}

func TestDB_CheckCookie(t *testing.T) {
	db := Init()
	cookie := http.Cookie{
		Name:  "cinsear_session",
		Value: "value",
		Path:  "/",
	}
	err := db.InsertCookie(cookie, 0)
	require.Nil(t, err)
	require.True(t, db.CheckCookie(cookie))
}

func TestDB_DeleteCookie(t *testing.T) {
	db := Init()
	ksyusha := models.NewUser{
		Credentials: models.Credentials{
			Email:    "ksyushag@mail.ru",
			Password: user.GetMD5Hash("qwerty123"),
		},
		Username: "Ksyusha",
	}
	e := db.Insert(ksyusha)
	require.Nil(t, e)
	cookie := http.Cookie{
		Name:  "cinsear_session",
		Value: "value",
		Path:  "/",
	}
	err := db.InsertCookie(cookie, 0)
	require.Nil(t, err)
	ok := db.CheckCookie(cookie)
	require.True(t, ok)
	db.DeleteCookie(cookie)
	ok = db.CheckCookie(cookie)
	require.False(t, ok)
}

func TestDB_FindFilmByTitle(t *testing.T) {
	db := Init()
	film := models.NewFilm{FilmInfo: models.FilmInfo{
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
	}}
	e := db.Insert(film)
	require.Nil(t, e)
	f, ok := db.FindFilmByTitle("Матрица")
	require.True(t, ok)
	require.Equal(t, f.Title, film.Title)
}

func TestDB_GetID(t *testing.T) {
	db := Init()
	ksyusha := models.NewUser{
		Credentials: models.Credentials{
			Email:    "ksyushag@mail.ru",
			Password: user.GetMD5Hash("qwerty123"),
		},
		Username: "Ksyusha",
	}
	e := db.Insert(ksyusha)
	require.Nil(t, e)
	id := db.GetID("users")
	require.Equal(t, 0, id)
}

func TestDB_FindByEmail(t *testing.T) {
	db := Init()
	ksyusha := models.NewUser{
		Credentials: models.Credentials{
			Email:    "ksyushag@mail.ru",
			Password: user.GetMD5Hash("qwerty123"),
		},
		Username: "Ksyusha",
	}
	e := db.Insert(ksyusha)
	require.Nil(t, e)
	u, ok := db.FindByEmail(ksyusha.Email)
	require.True(t, ok)
	require.Equal(t, ksyusha.Email, u.Email)
}

func TestDB_Get(t *testing.T) {
	db := Init()
	ksyusha := models.NewUser{
		Credentials: models.Credentials{
			Email:    "ksyushag@mail.ru",
			Password: user.GetMD5Hash("qwerty123"),
		},
		Username: "Ksyusha",
	}
	err := db.Insert(ksyusha)
	require.Nil(t, err)
	obj, err := db.Get(0, "user")
	require.Nil(t, err)
	u := obj.(models.User)
	require.True(t, reflect.DeepEqual(u.Credentials, ksyusha.Credentials))
}
