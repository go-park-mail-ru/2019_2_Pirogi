package domains

import (
	"github.com/asaskevich/govalidator"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/security"
	"html"
)

type FilmRepository interface {
	Insert(newFilm FilmNew) (ID, error)
	Update(id ID, film Film) error
	Delete(id ID) bool
	Get(id ID) Film
	GetMany(target Target, id ID) []Film
	MakeTrunc(film Film) FilmTrunc
	MakeFull(film Film) FilmFull
}

type FilmNew struct {
	Title       string   `json:"title" valid:"text, stringlength(1|50)"`
	Description string   `json:"description" valid:"text, stringlength(8|50)"`
	Year        int      `json:"year" valid:"year"`
	Countries   []string `json:"countries" valid:"countries"`
	Genres      []Genre  `json:"genres" valid:"genres"`
	PersonsID   []ID     `json:"persons_id" valid:"ids, optional"`
	Trailer     string   `json:"trailer" valid:"link, optional"`
}

func (fn *FilmNew) ToFilm(id ID) Film {
	return Film{
		ID:          id,
		Title:       html.EscapeString(fn.Title),
		Year:        fn.Year,
		Genres:      security.XSSFilterGenres(fn.Genres),
		Mark:        Mark(0),
		Description: html.EscapeString(fn.Description),
		Countries:   security.XSSFilterStrings(fn.Countries),
		PersonsID:   fn.PersonsID,
		Images:      []Image{"default.png"},
		ReviewsNum:  0,
		Trailer:     fn.Trailer,
	}
}

func (fn *FilmNew) Make(body []byte) error {
	err := fn.UnmarshalJSON(body)
	if err != nil {
		return err
	}
	_, err = govalidator.ValidateStruct(fn)
	return nil
}

type Film struct {
	ID          ID       `json:"id" bson:"_id" valid:"numeric,optional"`
	Title       string   `json:"title" valid:"text, stringlength(1|50)"`
	Year        int      `json:"year" valid:"year"`
	Genres      []Genre  `json:"genres" valid:"genres"`
	Mark        Mark     `json:"mark" valid:"mark"`
	Description string   `json:"description" valid:"text, stringlength(8|50)"`
	Countries   []string `json:"countries" valid:"countries"`
	PersonsID   []ID     `json:"persons_id" valid:"ids, optional"`
	Images      []Image  `json:"images" valid:"images, optional"`
	ReviewsNum  int      `json:"reviews_num" valid:"numeric, optional"`
	Trailer     string   `json:"trailer" valid:"link, optional"`
}

type FilmTrunc struct {
	ID     ID      `json:"id" valid:"numeric,optional"`
	Title  string  `json:"title" valid:"text, stringlength(1|50)"`
	Year   int     `json:"year" valid:"year"`
	Genres []Genre `json:"genres" valid:"genres"`
	Mark   Mark    `json:"mark" valid:"mark"`
	Image  Image   `json:"image" valid:"image"`
}

type FilmFull struct {
	ID          ID            `json:"id" bson:"_id" valid:"numeric,optional"`
	Title       string        `json:"title" valid:"text, stringlength(1|50)"`
	Year        int           `json:"year" valid:"year"`
	Genres      []Genre       `json:"genres" valid:"genres"`
	Mark        Mark          `json:"mark" valid:"mark"`
	Description string        `json:"description" valid:"text, stringlength(8|50)"`
	Countries   []string      `json:"countries" valid:"countries"`
	Persons     []PersonTrunc `json:"persons" valid:"optional"`
	Images      []Image       `json:"images" valid:"images, optional"`
	ReviewsNum  int           `json:"reviews_num" valid:"numeric, optional"`
	Trailer     string        `json:"trailer" valid:"link, optional"`
}

func (f *Film) SetMark(mark Mark) {
	f.Mark = mark
}

func (f *Film) Trunc() FilmTrunc {
	return FilmTrunc{
		ID:     f.ID,
		Title:  f.Title,
		Year:   f.Year,
		Genres: f.Genres,
		Mark:   f.Mark,
		Image:  f.Images[0],
	}
}

func (f *Film) Full(persons []Person) FilmFull {
	var personsTrunc []PersonTrunc
	for _, person := range persons {
		personsTrunc = append(personsTrunc, person.Trunc())
	}
	return FilmFull{
		ID:          f.ID,
		Title:       f.Title,
		Year:        f.Year,
		Genres:      f.Genres,
		Mark:        f.Mark,
		Description: f.Description,
		Countries:   f.Countries,
		Persons:     personsTrunc,
		Images:      f.Images,
		ReviewsNum:  f.ReviewsNum,
		Trailer:     f.Trailer,
	}
}
