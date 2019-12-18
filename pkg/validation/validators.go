package validation

import (
	"regexp"
	"strings"
	"time"

	valid "github.com/asaskevich/govalidator"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
)

//var datePattern = regexp.MustCompile("([0-9]{2}.){2}[1-2][0-9]{3}")
var imagePattern = regexp.MustCompile("(([0-9]|[a-z]){40}.(jpeg|jpg|png|gif))|default.png")
var textPattern = regexp.MustCompile(".+")
var linkPattern = regexp.MustCompile("http?://.+\\..*")

type validateFunc = func(i interface{}, o interface{}) bool

func InitValidator() {
	valid.SetFieldsRequiredByDefault(true)
	validators := map[string]validateFunc{
		"link": func(i interface{}, o interface{}) bool {
			subject, ok := i.(string)
			if !ok {
				return false
			}
			return linkPattern.MatchString(subject)
		},
		"ids": func(i interface{}, o interface{}) bool {
			subject, ok := i.([]model.ID)
			if !ok {
				return false
			}
			for _, id := range subject {
				if id < 0 {
					return false
				}
			}
			return true
		},
		"text": func(i interface{}, o interface{}) bool {
			subject, ok := i.(string)
			if !ok {
				return false
			}
			return textPattern.MatchString(subject)
		},
		"birthday": func(i interface{}, o interface{}) bool {
			subject, ok := i.(string)
			if !ok {
				return false
			}
			return len(strings.Split(subject, ".")) == 3
		},
		"mark": func(i interface{}, o interface{}) bool {
			subject, ok := i.(model.Mark)
			if !ok {
				return false
			}
			return subject >= 0.0 && subject <= 5.0
		},
		"countries": func(i interface{}, o interface{}) bool {
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
		},
		"genres": func(i interface{}, o interface{}) bool {
			subject, ok := i.([]model.Genre)
			if !ok {
				return false
			}
			for _, genre := range subject {
				if !validateGenre(string(genre)) {
					return false
				}
			}
			return true
		},
		"genre": func(i interface{}, o interface{}) bool {
			subject, ok := i.(string)
			if !ok {
				return false
			}
			return validateGenre(subject)
		},
		"target": func(i interface{}, o interface{}) bool {
			subject, ok := i.(model.Target)
			if !ok {
				return false
			}
			return validateTarget(string(subject))
		},
		"title": func(i interface{}, o interface{}) bool {
			subject, ok := i.(string)
			if !ok {
				return false
			}
			return len(subject) < 50
		},
		"description": func(i interface{}, o interface{}) bool {
			subject, ok := i.(string)
			if !ok {
				return false
			}
			return len(subject) < 500
		},
		"role": func(i interface{}, o interface{}) bool {
			subject, ok := i.(model.Target)
			if !ok {
				return false
			}
			return validateRole(string(subject))
		},
		"roles": func(i interface{}, o interface{}) bool {
			subject, ok := i.([]model.Role)
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
		},
		"time": func(i interface{}, o interface{}) bool {
			subject, ok := i.(time.Time)
			if !ok {
				return false
			}
			return subject.Year() <= time.Now().Year()
		},
		"image": func(i interface{}, o interface{}) bool {
			subject, ok := i.(model.Image)
			if !ok {
				return false
			}
			return imagePattern.MatchString(string(subject))
		},
		"images": func(i interface{}, o interface{}) bool {
			subject, ok := i.([]model.Image)
			if !ok {
				return false
			}
			for _, image := range subject {
				if !imagePattern.MatchString(string(image)) {
					return false
				}
			}
			return true
		},
		"password": func(i interface{}, o interface{}) bool {
			subject, ok := i.(string)
			if !ok {
				return false
			}
			return validatePassword(subject)
		},

		"year": func(i interface{}, o interface{}) bool {
			subject, ok := i.(int)
			if !ok {
				return false
			}
			return subject < 2020 && subject > 1800
		},
	}
	for key, value := range validators {
		valid.CustomTypeTagMap.Set(key, value)
	}
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

func validatePassword(password string) bool {
	return len(password) > 3
}
