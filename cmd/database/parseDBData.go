package main

import (
	"encoding/json"
	"errors"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/common"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"io"
	"log"
	"os"
	"path"
)

func parse(target string) ([]interface{}, error) {
	filename := configs.FileTargetMap[target]
	if filename == "" {
		return nil, errors.New("unsupported type")
	}
	reader, err := openFile(filename)
	if err != nil {
		return nil, err
	}
	dec := json.NewDecoder(reader)
	_, err = dec.Token()
	if err != nil {
		return nil, err
	}
	switch target {
	case configs.Default.PersonTargetName:
		var newPersons []models.NewPerson
		for dec.More() {
			var newPerson models.NewPerson
			err = dec.Decode(&newPerson)
			if err != nil {
				return nil, err
			}
			newPersons = append(newPersons, newPerson)
		}
		interfaces := make([]interface{}, len(newPersons))
		for i, val := range newPersons {
			interfaces[i] = val
		}
		return interfaces, nil
	case configs.Default.FilmTargetName:
		var newFilms []models.NewFilm
		for dec.More() {
			var newFilm models.NewFilm
			err = dec.Decode(&newFilm)
			if err != nil {
				return nil, err
			}
			newFilms = append(newFilms, newFilm)
		}
		interfaces := make([]interface{}, len(newFilms))
		for i, val := range newFilms {
			interfaces[i] = val
		}
		return interfaces, nil
	case configs.Default.ReviewTargetName:
		var newReviews []models.NewReview
		for dec.More() {
			var newReview models.NewReview
			err = dec.Decode(&newReview)
			if err != nil {
				return nil, err
			}
			newReviews = append(newReviews, newReview)
		}
		interfaces := make([]interface{}, len(newReviews))
		for i, val := range newReviews {
			interfaces[i] = val
		}
		return interfaces, nil
	case configs.Default.StarTargetName:
		var newStars []models.Stars
		for dec.More() {
			var newStar models.Stars
			err = dec.Decode(&newStar)
			if err != nil {
				return nil, err
			}
			newStars = append(newStars, newStar)
		}
		interfaces := make([]interface{}, len(newStars))
		for i, val := range newStars {
			interfaces[i] = val
		}
		return interfaces, nil
	case configs.Default.UserTargetName:
		var newUsers []models.NewUser
		for dec.More() {
			var newUser models.NewUser
			err = dec.Decode(&newUser)
			if err != nil {
				return nil, err
			}
			newUsers = append(newUsers, newUser)
		}
		interfaces := make([]interface{}, len(newUsers))
		for i, val := range newUsers {
			interfaces[i] = val
		}
		return interfaces, nil
	default:
		return nil, errors.New("unsupported target")
	}
}

func openFile(filename string) (io.ReadCloser, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	file, err := os.Open(path.Join(wd, "/cmd/database", filename))
	if err != nil {
		return nil, err
	}
	reader := io.ReadCloser(file)
	return reader, nil
}

func main() {
	configPath := "configs"
	err := common.UnmarshalConfigs(&configPath)
	if err != nil {
		log.Fatal(err)
	}
}
