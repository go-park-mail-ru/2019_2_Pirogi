package makers

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/domains"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

const imageFilename = "9d16a00dcbc3778f4e48962c3b8c8f0b4d662410.png"

var personNew = domains.NewPerson{
	Name:       personFull.Name,
	Roles:      personFull.Roles,
	Birthday:   personFull.Birthday,
	Birthplace: personFull.Birthplace,
}

var personFull = domains.PersonFull{
	ID:         2,
	Name:       "Artyom",
	Mark:       domains.Mark(3.5),
	Roles:      []domains.Role{"actor", "director"},
	Birthday:   "09.12.1998",
	Birthplace: "Russia",
	Genres:     []domains.Genre{"трагикомедия"},
	Films:      []domains.FilmTrunc{filmTrunc},
	Likes:      34,
	Images:     []domains.Image{imageFilename, imageFilename},
}

var person = domains.Person{
	ID:         personFull.ID,
	Name:       personFull.Name,
	Mark:       personFull.Mark,
	Roles:      personFull.Roles,
	Birthday:   personFull.Birthday,
	Birthplace: personFull.Birthplace,
	Genres:     personFull.Genres,
	FilmsID:    []domains.ID{2},
	Likes:      personFull.Likes,
	Images:     personFull.Images,
}

var personTrunc = domains.PersonTrunc{
	ID:   2,
	Name: "Artyom",
}

var filmNew = domains.NewFilm{
	Title:       filmFull.Title,
	Description: filmFull.Title,
	Year:        filmFull.Year,
	Countries:   filmFull.Countries,
	Genres:      filmFull.Genres,
	PersonsID:   film.PersonsID,
}

var filmFull = domains.FilmFull{
	ID:          2,
	Title:       "Matrix",
	Year:        "1998",
	Genres:      []domains.Genre{"драма"},
	Mark:        domains.Mark(3.5),
	Description: "film about matrix",
	Countries:   []string{"USA", "Russia"},
	Persons:     []domains.PersonTrunc{personTrunc},
	Images:      []domains.Image{imageFilename, imageFilename},
	ReviewsNum:  5,
}

var film = domains.Film{
	ID:          filmFull.ID,
	Title:       filmFull.Title,
	Year:        filmFull.Year,
	Genres:      filmFull.Genres,
	Mark:        filmFull.Mark,
	Description: filmFull.Description,
	Countries:   filmFull.Countries,
	PersonsID:   []domains.ID{2},
	Images:      filmFull.Images,
	ReviewsNum:  5,
}

var filmTrunc = domains.FilmTrunc{
	ID:     2,
	Title:  "Matrix",
	Year:   "1998",
	Genres: []domains.Genre{"драма"},
	Mark:   domains.Mark(3.5),
}

var reviewNew = domains.NewReview{
	Title:    review.Title,
	Body:     review.Body,
	FilmID:   review.FilmID,
	AuthorID: review.AuthorID,
}

var review = domains.Review{
	ID:       2,
	Title:    "title",
	Body:     "body",
	FilmID:   2,
	AuthorID: 2,
	Date:     time.Time{},
	Likes:    0,
}

func TestMakeTruncFilm(t *testing.T) {
	expected := filmTrunc
	actual := MakeFilmTrunc(film)
	require.Equal(t, expected, actual)
}

func TestMakeFilm(t *testing.T) {
	expected := domains.Film{
		ID:          2,
		Title:       filmNew.Title,
		Year:        filmNew.Year,
		Genres:      filmNew.Genres,
		Mark:        domains.Mark(0),
		Description: filmNew.Description,
		Countries:   filmNew.Countries,
		PersonsID:   filmNew.PersonsID,
		Images:      []domains.Image{"default.png"},
		ReviewsNum:  0,
	}
	actual := MakeFilm(2, &filmNew)
	require.Equal(t, expected, actual)
}

func TestMakeFullFilm(t *testing.T) {
	expected := filmFull
	actual := MakeFilmFull(film, []domains.Person{person})
	require.Equal(t, expected, actual)
}

func TestMakeTruncPerson(t *testing.T) {
	expected := personTrunc
	actual := MakeTruncPerson(person)
	require.Equal(t, expected, actual)
}

func TestMakePerson(t *testing.T) {
	expected := domains.Person{
		ID:         person.ID,
		Name:       person.Name,
		Mark:       0,
		Roles:      person.Roles,
		Birthday:   personFull.Birthday,
		Birthplace: personFull.Birthplace,
		Genres:     []domains.Genre{},
		FilmsID:    []domains.ID{},
		Likes:      0,
		Images:     []domains.Image{"default.png"},
	}
	actual := MakePerson(2, personNew)
	require.Equal(t, expected, actual)
}

func TestMakeFullPerson(t *testing.T) {
	expected := personFull
	actual := MakeFullPerson(person, []domains.Film{film})
	require.Equal(t, expected, actual)
}

func TestMakeReview(t *testing.T) {
	expected := review
	actual := MakeReview(2, reviewNew)
	expected.Date = actual.Date
	require.Equal(t, expected, actual)
}
