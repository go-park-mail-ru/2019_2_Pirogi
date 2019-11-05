package common

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"html"
)

func XSSFilterRoles(roles []models.Role) []models.Role {
	for _, role := range roles {
		role = models.Role(html.EscapeString(string(role)))
	}
	return roles
}

func XSSFilterGenres(Genres []models.Genre) []models.Genre {
	for _, Genre := range Genres {
		Genre = models.Genre(html.EscapeString(string(Genre)))
	}
	return Genres
}

func XSSFilterStrings(in []string) []string {
	for _, val := range in {
		val = html.EscapeString(val)
	}
	return in
}
