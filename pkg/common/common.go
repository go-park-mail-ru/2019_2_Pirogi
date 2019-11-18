package common

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/domains/film"
)

func TruncFilms(in []film.Film) (out []film.FilmTrunc) {
	for _, film := range in {
		out = append(out, film.Trunc())
	}
	return out
}
