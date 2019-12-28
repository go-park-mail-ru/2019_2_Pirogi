package fixture

import "github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"

var FilmID = model.ID(1)
var Title = "test film"
var Year = 1998
var Countries = []string{"Russia", "USA"}
var Genres = []model.Genre{"драма", "ужасы"}
var PersonsID = []model.ID{1, 2, 3}
var Trailer = "yXzVL_srdU0"
var Mark = model.Mark(4.3)
var Stars = model.Stars(3)
var Films = []model.Film{Film, Film}

var FilmNew = model.FilmNew{
	Title:       Title,
	Description: Description,
	Year:        Year,
	Countries:   Countries,
	Genres:      Genres,
	PersonsID:   PersonsID,
	Trailer:     Trailer,
}

var Film = model.Film{
	ID:          FilmID,
	Title:       Title,
	Year:        Year,
	Genres:      Genres,
	Mark:        Mark,
	Description: Description,
	Countries:   Countries,
	PersonsID:   PersonsID,
	Images:      []model.Image{Image},
	ReviewsNum:  0,
	Trailer:     Trailer,
}

var FilmTrunc = Film.Trunc()
var FilmFull = Film.Full(Persons)
