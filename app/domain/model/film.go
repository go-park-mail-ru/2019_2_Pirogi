package model

import "github.com/go-park-mail-ru/2019_2_Pirogi/configs"

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
		Title:       fn.Title,
		Year:        fn.Year,
		Genres:      fn.Genres,
		Mark:        Mark(0),
		Description: fn.Description,
		Countries:   fn.Countries,
		PersonsID:   fn.PersonsID,
		Images:      []Image{Image(configs.Default.DefaultImageName)},
		ReviewsNum:  0,
		Trailer:     fn.Trailer,
	}
}

func (fn *FilmNew) Make(body []byte) error {
	err := fn.UnmarshalJSON(body)
	if err != nil {
		return err
	}
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
	Trailer     string   `json:"trailer" valid:"text, optional"`
}

type FilmTrunc struct {
	ID          ID            `json:"id" valid:"numeric,optional"`
	Title       string        `json:"title" valid:"text, stringlength(1|50)"`
	Year        int           `json:"year" valid:"year"`
	Genres      []Genre       `json:"genres" valid:"genres"`
	Mark        Mark          `json:"mark" valid:"mark"`
	Description string        `json:"description" valid:"text, stringlength(8|50)"`
	Persons     []PersonTrunc `json:"persons" valid:"optional"`
	Image       Image         `json:"image" valid:"image"`
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
	Trailer     string        `json:"trailer" valid:"text, optional"`
	Related     []FilmTrunc   `json:"related,omitempty" valid:"optional"`
}

func (f *Film) SetMark(mark Mark) {
	f.Mark = mark
}

func (f *Film) Trunc() FilmTrunc {
	return FilmTrunc{
		ID:          f.ID,
		Title:       f.Title,
		Year:        f.Year,
		Description: f.Description,
		Genres:      f.Genres,
		Mark:        f.Mark,
		Image:       f.Images[0],
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
