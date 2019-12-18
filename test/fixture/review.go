package fixture

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"time"
)

var ReviewID = model.ID(1)
var ReviewTitle = "test review"
var ReviewBody = "test review body"
var Date = time.Date(2018, time.December, 10, 23, 0, 0, 0, time.UTC)

var Review = model.Review{
	ID:       ReviewID,
	Title:    ReviewTitle,
	Body:     ReviewBody,
	FilmID:   FilmID,
	AuthorID: UserID,
	Date:     Date,
	Mark:     Mark,
}
