package validators

import (
	valid "github.com/asaskevich/govalidator"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"regexp"
	"time"
)

var yearPattern = regexp.MustCompile("[1|2][0-9]{3}")
var datePattern = regexp.MustCompile("([0-9]{2}.){2}[1-2][0-9]{3}")
var imagePattern = regexp.MustCompile("([0-9]|[a-z]){40}.(jpeg|jpg|png|gif)")

func InitValidator() {
	valid.SetFieldsRequiredByDefault(true)

	valid.CustomTypeTagMap.Set("year", func(i interface{}, o interface{}) bool {
		subject, ok := i.(string)
		if !ok {
			return false
		}
		return yearPattern.MatchString(subject)
	})

	valid.CustomTypeTagMap.Set("date", func(i interface{}, o interface{}) bool {
		subject, ok := i.(string)
		if !ok {
			return false
		}
		return datePattern.MatchString(subject)
	})

	valid.CustomTypeTagMap.Set("mark", func(i interface{}, o interface{}) bool {
		subject, ok := i.(models.Mark)
		if !ok {
			return false
		}
		return subject >= 0.0 && subject <= 5.0
	})

	valid.CustomTypeTagMap.Set("countries", func(i interface{}, o interface{}) bool {
		subject, ok := i.([]string)
		if !ok {
			return false
		}
		for _, val := range subject {
			if !validateCountry(val) {
				return false
			}
		}
		return true
	})

	valid.CustomTypeTagMap.Set("genres", func(i interface{}, o interface{}) bool {
		subject, ok := i.([]models.Genre)
		if !ok {
			return false
		}
		for _, genre := range subject {
			if !validateGenre(string(genre)) {
				return false
			}
		}
		return true
	})

	valid.CustomTypeTagMap.Set("target", func(i interface{}, o interface{}) bool {
		subject, ok := i.(models.Target)
		if !ok {
			return false
		}
		return validateTarget(string(subject))
	})

	valid.CustomTypeTagMap.Set("title", func(i interface{}, o interface{}) bool {
		subject, ok := i.(string)
		if !ok {
			return false
		}
		return len(subject) < 50
	})

	valid.CustomTypeTagMap.Set("description", func(i interface{}, o interface{}) bool {
		subject, ok := i.(string)
		if !ok {
			return false
		}
		return len(subject) < 500
	})

	valid.CustomTypeTagMap.Set("role", func(i interface{}, o interface{}) bool {
		subject, ok := i.(models.Target)
		if !ok {
			return false
		}
		return validateRole(string(subject))
	})

	valid.CustomTypeTagMap.Set("roles", func(i interface{}, o interface{}) bool {
		subject, ok := i.([]models.Role)
		if !ok {
			return false
		}
		for _, val := range subject {
			ok := validateRole(string(val))
			if !ok {
				return false
			}
		}
		return true
	})

	valid.CustomTypeTagMap.Set("time", func(i interface{}, o interface{}) bool {
		subject, ok := i.(time.Time)
		if !ok {
			return false
		}
		return subject.Year() <= time.Now().Year()
	})

	valid.CustomTypeTagMap.Set("image", func(i interface{}, o interface{}) bool {
		subject, ok := i.(string)
		if !ok {
			return false
		}
		return imagePattern.MatchString(subject)
	})

	valid.CustomTypeTagMap.Set("films_trunc", func(i interface{}, o interface{}) bool {
		subject, ok := i.([]models.FilmTrunc)
		if !ok {
			return false
		}
		for _, filmT := range subject {
			ok, _ := valid.ValidateStruct(filmT)
			if !ok {
				return false
			}
		}
		return true
	})

	valid.CustomTypeTagMap.Set("persons_trunc", func(i interface{}, o interface{}) bool {
		subject, ok := i.([]models.PersonTrunc)
		if !ok {
			return false
		}
		for _, personT := range subject {
			ok, _ := valid.ValidateStruct(personT)
			if !ok {
				return false
			}
		}
		return true
	})

	valid.CustomTypeTagMap.Set("images", func(i interface{}, o interface{}) bool {
		subject, ok := i.([]models.Image)
		if !ok {
			return false
		}
		for _, image := range subject {
			ok, _ := valid.ValidateStruct(image)
			if !ok {
				return false
			}
		}
		return true
	})

}

func validateCountry(country string) bool {
	if !valid.IsAlpha(country) {
		return false
	}
	return true
}

func validateGenre(genre string) bool {
	for _, confGenre := range configs.Genres {
		if genre == confGenre {
			return true
		}
	}
	return false
}

func validateTarget(target string) bool {
	for _, confTarget := range configs.Targets {
		if target == confTarget {
			return true
		}
	}
	return false
}

func validateRole(role string) bool {
	for _, confRole := range configs.Roles {
		if role == confRole {
			return true
		}
	}
	return false
}
