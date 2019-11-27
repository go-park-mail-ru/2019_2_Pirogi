package security

import (
	"html"

	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
)

func XSSFilterRoles(roles []model.Role) []model.Role {
	for idx := range roles {
		roles[idx] = model.Role(html.EscapeString(string(roles[idx])))
	}
	return roles
}

func XSSFilterGenres(genres []model.Genre) []model.Genre {
	for idx := range genres {
		genres[idx] = model.Genre(html.EscapeString(string(genres[idx])))
	}
	return genres
}

func XSSFilterStrings(in []string) []string {
	for idx := range in {
		in[idx] = html.EscapeString(in[idx])
	}
	return in
}
