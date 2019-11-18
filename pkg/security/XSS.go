package security

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/domains"
	"html"
)

func XSSFilterRoles(roles []domains.Role) []domains.Role {
	for idx := range roles {
		roles[idx] = domains.Role(html.EscapeString(string(roles[idx])))
	}
	return roles
}

func XSSFilterGenres(genres []domains.Genre) []domains.Genre {
	for idx := range genres {
		genres[idx] = domains.Genre(html.EscapeString(string(genres[idx])))
	}
	return genres
}

func XSSFilterStrings(in []string) []string {
	for idx := range in {
		in[idx] = html.EscapeString(in[idx])
	}
	return in
}
