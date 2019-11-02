package validators

import (
	"testing"
	"time"

	valid "github.com/asaskevich/govalidator"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/stretchr/testify/require"
)

const imageFilename = "9d16a00dcbc3778f4e48962c3b8c8f0b4d662410.png"

func TestLike(t *testing.T) {
	InitValidator()
	like := models.Like{
		UserID:   2,
		Target:   models.Target("film"),
		TargetID: 5,
	}
	_, err := valid.ValidateStruct(like)
	require.NoError(t, err)
}

func TestStars(t *testing.T) {
	InitValidator()
	stars := models.Stars{
		UserID: 3,
		FilmID: 4,
		Mark:   4.2,
	}
	_, err := valid.ValidateStruct(stars)
	require.NoError(t, err)
}

func TestImage(t *testing.T) {
	InitValidator()
	image := models.Image{
		ID:       12,
		Filename: "cffa50a32cb13a240d705317bcec65dd1f31b6ad.jpg",
	}
	_, err := valid.ValidateStruct(image)
	require.NoError(t, err)
}

func TestTruncUser(t *testing.T) {
	InitValidator()
	userTrunc := models.UserTrunc{
		ID:          2,
		Username:    "Artefakt1",
		Mark:        3.5,
		Description: "lalalalal",
		Image: models.Image{
			ID:       2,
			Filename: "cffa50a32cb13a240d705317bcec65dd1f31b6ad.jpg",
		},
	}
	_, err := valid.ValidateStruct(userTrunc)
	require.NoError(t, err)
}

func TestTruncFilm(t *testing.T) {
	filmTrunc := models.FilmTrunc{
		ID:     2,
		Title:  "matrix",
		Year:   "1998",
		Genres: []models.Genre{"драма"},
		Poster: models.Image{
			ID:       12,
			Filename: imageFilename,
		},
		Mark: 3.4,
	}
	_, err := valid.ValidateStruct(filmTrunc)
	require.NoError(t, err)
}

func TestCredentials(t *testing.T) {
	credentials := models.Credentials{
		Email:    "bakulev.artyom@artbakulev.com",
		Password: "qwertyiop12",
	}
	_, err := valid.ValidateStruct(credentials)
	require.NoError(t, err)
}

func TestPersonTrunc(t *testing.T) {
	InitValidator()
	personTrunc := models.PersonTrunc{
		ID:   12,
		Name: "artefakt",
		Mark: models.Mark(0.4),
	}
	_, err := valid.ValidateStruct(personTrunc)
	require.NoError(t, err)
}

func TestPerson(t *testing.T) {
	InitValidator()
	image := models.Image{
		ID:       12,
		Filename: "cffa50a32cb13a240d705317bcec65dd1f31b6ad.jpg",
	}
	personTrunc := models.PersonTrunc{
		ID:   12,
		Name: "artefakt",
		Mark: models.Mark(0.4),
	}
	filmTrunc := models.FilmTrunc{
		ID:     2,
		Title:  "matrix",
		Year:   "1998",
		Genres: []models.Genre{"драма"},
		Poster: models.Image{
			ID:       12,
			Filename: imageFilename,
		},
		Mark: 3.0,
	}
	person := models.Person{
		PersonTrunc: personTrunc,
		Roles:       []models.Role{"actor"},
		Birthday:    "09.12.1998",
		Birthplace:  "USA",
		Genres:      []models.Genre{"драма"},
		FilmsID: []models.FilmTrunc{
			filmTrunc,
		},
		Likes: 2,
		ImagesID: []models.Image{
			image,
		},
	}
	_, err := valid.ValidateStruct(person)
	require.NoError(t, err)
}

func TestFilm(t *testing.T) {
	image := models.Image{
		ID:       12,
		Filename: "cffa50a32cb13a240d705317bcec65dd1f31b6ad.jpg",
	}
	personTrunc := models.PersonTrunc{
		ID:   12,
		Name: "artefakt",
		Mark: models.Mark(0.4),
	}
	filmTrunc := models.FilmTrunc{
		ID:     2,
		Title:  "matrix",
		Year:   "1998",
		Genres: []models.Genre{"драма"},
		Poster: models.Image{
			ID:       12,
			Filename: imageFilename,
		},
		Mark: 3.0,
	}
	film := models.Film{
		FilmTrunc:   filmTrunc,
		Countries:   []string{"USA"},
		Description: "laslasldasldlasdl",
		PersonsID: []models.PersonTrunc{
			personTrunc,
		},
		Directors: []models.PersonTrunc{
			personTrunc,
		},
		Images: []models.Image{
			image,
		},
		ReviewsNum: 23,
	}
	_, err := valid.ValidateStruct(film)
	require.NoError(t, err)
}
func TestReview(t *testing.T) {
	InitValidator()
	review := models.Review{
		NewReview: models.NewReview{
			Title:    "Обычный обзор",
			Body:     "лалааллфылвфлывлфывлфывйцуasd12",
			FilmID:   2,
			AuthorID: 6,
		},
		ID:    3,
		Date:  time.Now(),
		Likes: 8,
	}
	_, err := valid.ValidateStruct(review)
	require.NoError(t, err)
}
