package validation

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"testing"
	"time"

	valid "github.com/asaskevich/govalidator"
	"github.com/stretchr/testify/require"
)

func TestTarget(t *testing.T) {
	InitValidator()
	target := model.Target("film")
	ok := validateTarget(string(target))
	require.True(t, ok)
}

func TestLike(t *testing.T) {
	InitValidator()
	like := model.Like{
		UserID:   2,
		Target:   model.Target("film"),
		TargetID: 5,
	}
	_, err := valid.ValidateStruct(like)
	require.NoError(t, err)
}

func TestStars(t *testing.T) {
	InitValidator()
	stars := model.Stars{
		UserID: 3,
		FilmID: 4,
		Mark:   4.2,
	}
	_, err := valid.ValidateStruct(stars)
	require.NoError(t, err)
}

func TestTruncUser(t *testing.T) {
	InitValidator()
	userTrunc := model.UserTrunc{
		ID:          2,
		Username:    "Artefakt1",
		Mark:        3.5,
		Description: "lalalalal",
		Image:       "cffa50a32cb13a240d705317bcec65dd1f31b6ad.jpg",
		Reviews:     0,
	}
	_, err := valid.ValidateStruct(userTrunc)
	require.NoError(t, err)
}

func TestTruncFilm(t *testing.T) {
	filmTrunc := model.FilmTrunc{
		ID:     2,
		Title:  "matrix",
		Year:   1998,
		Genres: []model.Genre{"драма"},
		Mark:   3.4,
	}
	_, err := valid.ValidateStruct(filmTrunc)
	require.NoError(t, err)
}

func TestCredentials(t *testing.T) {
	credentials := model.UserCredentials{
		Email:    "bakulev.artyom@artbakulev.com",
		Password: "qwertyiop12",
	}
	_, err := valid.ValidateStruct(credentials)
	require.NoError(t, err)
}

func TestPersonTrunc(t *testing.T) {
	InitValidator()
	personTrunc := model.PersonTrunc{
		ID:   12,
		Name: "artefakt",
	}
	_, err := valid.ValidateStruct(personTrunc)
	require.NoError(t, err)
}

func TestPerson(t *testing.T) {
	InitValidator()
	image := model.Image("cffa50a32cb13a240d705317bcec65dd1f31b6ad.jpg")

	personTrunc := model.PersonTrunc{
		ID:   12,
		Name: "artefakt",
	}
	person := model.Person{
		ID:         personTrunc.ID,
		Name:       personTrunc.Name,
		Mark:       model.Mark(2.3),
		Roles:      []model.Role{"actor"},
		Birthday:   "09.12.1998",
		Birthplace: "USA",
		Genres:     []model.Genre{"драма"},
		FilmsID: []model.ID{
			0,
		},
		Likes: 2,
		Images: []model.Image{
			image,
		},
	}
	_, err := valid.ValidateStruct(person)
	require.NoError(t, err)
}

func TestFilm(t *testing.T) {
	image := model.Image("cffa50a32cb13a240d705317bcec65dd1f31b6ad.jpg")
	filmTrunc := model.FilmTrunc{
		ID:     2,
		Title:  "matrix",
		Year:   1998,
		Genres: []model.Genre{"драма"},
		Mark:   3.0,
	}
	film := model.Film{
		ID:          filmTrunc.ID,
		Title:       filmTrunc.Title,
		Year:        filmTrunc.Year,
		Genres:      filmTrunc.Genres,
		Mark:        filmTrunc.Mark,
		Countries:   []string{"USA"},
		Description: "laslasldasldlasdl",
		PersonsID: []model.ID{
			0,
		},
		Images: []model.Image{
			image,
		},
		ReviewsNum: 23,
	}
	_, err := valid.ValidateStruct(film)
	require.NoError(t, err)
}
func TestReview(t *testing.T) {
	InitValidator()
	review := model.Review{
		Title:    "Обычный обзор",
		Body:     "лалааллфылвфлывлфывлфывйцуasd12",
		FilmID:   2,
		AuthorID: 6,
		ID:       3,
		Date:     time.Now(),
		Likes:    8,
	}
	_, err := valid.ValidateStruct(review)
	require.NoError(t, err)
}
