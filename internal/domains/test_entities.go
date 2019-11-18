package domains

var testFilmNew = FilmNew{
	Title:       "Matrix",
	Description: "Fantasy",
	Year:        1998,
	Countries:   []string{"USA"},
	Genres:      []Genre{"Драма"},
	PersonsID:   []ID{1, 2, 3},
}

var testFilm = testFilmNew.ToFilm(2)
var testFilmTrunc = testFilm.Trunc()
var testFilmFull = testFilm.Full([]Person{testPerson})

var testPersonNew = PersonNew{
	Name:       "Artyom",
	Roles:      []Role{"actor"},
	Birthday:   "09.12.1998",
	Birthplace: "Russia",
}

var testPerson = testPersonNew.ToPerson(2)
var testPersonTrunc = testPerson.Trunc()
var testPersonFull = testPerson.Full([]Film{testFilm, testFilm})

var testReviewNew = ReviewNew{
	Title:    "title",
	Body:     "body",
	FilmID:   2,
	AuthorID: 2,
}

var testReview = testReviewNew.ToReview(2)
