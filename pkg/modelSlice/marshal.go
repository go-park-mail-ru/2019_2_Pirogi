package modelSlice

import "github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"

func MarshalFilmsTrunc(filmsTrunc []model.FilmTrunc) (body [][]byte) {
	for _, filmTrunc := range filmsTrunc {
		raw, err := filmTrunc.MarshalJSON()
		if err != nil {
			continue
		}
		body = append(body, raw)
	}
	return body
}

func MarshalPersonsTrunc(personsTrunc []model.PersonTrunc) (body [][]byte) {
	for _, personTrunc := range personsTrunc {
		raw, err := personTrunc.MarshalJSON()
		if err != nil {
			continue
		}
		body = append(body, raw)
	}
	return body
}
