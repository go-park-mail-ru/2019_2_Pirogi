package security

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"html"
)

func XSSFilterRoles(roles []models.Role) []models.Role {
	for idx := range roles {
		roles[idx] = models.Role(html.EscapeString(string(roles[idx])))
	}
	return roles
}

func XSSFilterGenres(genres []models.Genre) []models.Genre {
	for idx := range genres {
		genres[idx] = models.Genre(html.EscapeString(string(genres[idx])))
	}
	return genres
}

func XSSFilterStrings(in []string) []string {
	for idx := range in {
		in[idx] = html.EscapeString(in[idx])
	}
	return in
}
