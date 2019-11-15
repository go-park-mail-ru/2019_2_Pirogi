package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/common"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/database"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/labstack/gommon/log"
	"io"
	"os"
	"path"
)

func parse(target string) ([]interface{}, error) {
	filename := configs.FileTargetMap[target]
	if filename == "" {
		return nil, errors.New("unsupported type")
	}
	readCloser, err := openFile(filename)
	if err != nil {
		return nil, err
	}
	defer readCloser.Close()
	dec := json.NewDecoder(readCloser)
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

func FakeFillDB(conn *database.MongoConnection) {
	persons, err := parse(configs.Default.PersonTargetName)
	if err != nil {
		log.Fatal(err)
	}
	for i, person := range persons {
		fmt.Printf("inserting %d person from %d\n", i, len(persons))
		conn.Upsert(person)
	}

	films, err := parse(configs.Default.FilmTargetName)
	if err != nil {
		log.Fatal(err)
	}
	for i, film := range films {
		fmt.Printf("inserting %d film from %d\n", i, len(films))
		conn.Upsert(film)
		newFilm := film.(models.NewFilm)
		for _, id := range newFilm.PersonsID {
			data, err := conn.Get(id, configs.Default.PersonTargetName)
			if err != nil {
				continue
			}
			person := data.(models.Person)
			person.FilmsID = append(person.FilmsID, models.ID(i))
			conn.Upsert(person)
		}

	}
}

func main() {
	configsPath := flag.String("config", "configs/", "directory with configs")
	flag.Parse()

	err := common.UnmarshalConfigs(configsPath)
	if err != nil {
		log.Fatal(err.Error())
	}

	conn, err := database.InitMongo("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}

	conn.ClearDB()
	err = conn.InitCounters()
	if err != nil {
		log.Fatal(err)
	}
	FakeFillDB(conn)
}
