package validators

import (
	valid "github.com/asaskevich/govalidator"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/domains"
	"regexp"
	"time"
)

var datePattern = regexp.MustCompile("([0-9]{2}.){2}[1-2][0-9]{3}")
var imagePattern = regexp.MustCompile("(([0-9]|[a-z]){40}.(jpeg|jpg|png|gif))|default.png")
var textPattern = regexp.MustCompile(".+")
var linkPattern = regexp.MustCompile("http?://.+\\..*")

func InitValidator() {
	valid.SetFieldsRequiredByDefault(true)

	valid.CustomTypeTagMap.Set("link", func(i interface{}, o interface{}) bool {
		subject, ok := i.(string)
		if !ok {
			return false
		}
		return linkPattern.MatchString(subject)
	})
	valid.CustomTypeTagMap.Set("ids", func(i interface{}, o interface{}) bool {
		subject, ok := i.([]domains.ID)
		if !ok {
			return false
		}
		for _, id := range subject {
			if id < 0 {
				return false
			}
		}
		return true
	})

	valid.CustomTypeTagMap.Set("text", func(i interface{}, o interface{}) bool {
		subject, ok := i.(string)
		if !ok {
			return false
		}
		return textPattern.MatchString(subject)
	})
	valid.CustomTypeTagMap.Set("year", func(i interface{}, o interface{}) bool {
		subject, ok := i.(int)
		if !ok {
			return false
		}
		return subject > 0
	})

	valid.CustomTypeTagMap.Set("date", func(i interface{}, o interface{}) bool {
		subject, ok := i.(string)
		if !ok {
			return false
		}
		return datePattern.MatchString(subject)
	})

	valid.CustomTypeTagMap.Set("mark", func(i interface{}, o interface{}) bool {
		subject, ok := i.(domains.Mark)
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
		subject, ok := i.([]domains.Genre)
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

	valid.CustomTypeTagMap.Set("genre", func(i interface{}, o interface{}) bool {
		subject, ok := i.(string)
		if !ok {
			return false
		}
		return validateGenre(subject)
	})

	valid.CustomTypeTagMap.Set("target", func(i interface{}, o interface{}) bool {
		subject, ok := i.(domains.Target)
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
		subject, ok := i.(domains.Target)
		if !ok {
			return false
		}
		return validateRole(string(subject))
	})

	valid.CustomTypeTagMap.Set("roles", func(i interface{}, o interface{}) bool {
		subject, ok := i.([]domains.Role)
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
		subject, ok := i.(domains.Image)
		if !ok {
			return false
		}
		return imagePattern.MatchString(string(subject))
	})

	valid.CustomTypeTagMap.Set("images", func(i interface{}, o interface{}) bool {
		subject, ok := i.([]domains.Image)
		if !ok {
			return false
		}
		for _, image := range subject {
			if !imagePattern.MatchString(string(image)) {
				return false
			}
		}
		return true
	})

	valid.CustomTypeTagMap.Set("password", func(i interface{}, o interface{}) bool {
		subject, ok := i.(string)
		if !ok {
			return false
		}
		return validatePassword(subject)
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

//TODO: в продакшене расскоментировать
func validatePassword(password string) bool {
	return len(password) > 3
	//letters := 0
	//var flags = []bool{false, false, false, false}
	//for _, c := range password {
	//	switch {
	//case unicode.IsNumber(c):
	//	flags[0] = true
	//case unicode.IsUpper(c):
	//	flags[1] = true
	//	letters++
	//case unicode.IsPunct(c) || unicode.IsSymbol(c):
	//	flags[2] = true
	//	case unicode.IsLetter(c) || c == ' ':
	//		letters++
	//	default:
	//		return false
	//	}
	//
	//}
	//flags[3] = letters > 7
	//if flags[0] && flags[1] && flags[2] && flags[3] {
	//	return true
	//}
	//return letters > 7
}
