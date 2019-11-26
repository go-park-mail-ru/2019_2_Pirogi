package model_tests

import "github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"

var testFilmNew = model.FilmNew{
	Title:       "Matrix",
	Description: "Fantasy",
	Year:        1998,
	Countries:   []string{"USA"},
	Genres:      []model.Genre{"Драма"},
	PersonsID:   []model.ID{1, 2, 3},
}

var testFilm = testFilmNew.ToFilm(2)
var testFilmTrunc = testFilm.Trunc()
var testFilmFull = testFilm.Full([]model.Person{testPerson})

var testPersonNew = model.PersonNew{
	Name:       "Artyom",
	Roles:      []model.Role{"actor"},
	Birthday:   "09.12.1998",
	Birthplace: "Russia",
}

var testPerson = testPersonNew.ToPerson(2)
var testPersonTrunc = testPerson.Trunc()
var testPersonFull = testPerson.Full([]model.Film{testFilm, testFilm})

var testReviewNew = model.ReviewNew{
	Title:    "title",
	Body:     "body",
	FilmID:   2,
	AuthorID: 2,
}

var testReview = testReviewNew.ToReview(2)
