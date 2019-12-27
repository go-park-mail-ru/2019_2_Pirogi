package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/infrastructure/database"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/configuration"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/files"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/hash"
	"github.com/labstack/gommon/log"
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
		var newPersons []model.PersonNew
		for dec.More() {
			var newPerson model.PersonNew
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
		var newFilms []model.FilmNew
		for dec.More() {
			var newFilm model.FilmNew
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
		var newReviews []model.ReviewNew
		for dec.More() {
			var newReview model.ReviewNew
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
		var newRatings []model.Rating
		for dec.More() {
			var newRating model.Rating
			err = dec.Decode(&newRating)
			if err != nil {
				return nil, err
			}
			newRatings = append(newRatings, newRating)
		}
		interfaces := make([]interface{}, len(newRatings))
		for i, val := range newRatings {
			interfaces[i] = val
		}
		return interfaces, nil
	case configs.Default.UserTargetName:
		var newUsers []model.UserNew
		for dec.More() {
			var newUser model.UserNew
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
	case configs.Default.FilmImageTargetName, configs.Default.PersonImageTargetName:
		var images []model.Image
		for dec.More() {
			var image model.Image
			err = dec.Decode(&image)
			if err != nil {
				return nil, err
			}
			images = append(images, image)
		}
		interfaces := make([]interface{}, len(images))
		for i, val := range images {
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

func uploadAndSaveImage(url string, baseFolder string) (string, error) {
	//TODO: переделать)))
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	ending, e := files.DetectContentType(body)
	if e != nil {
		return "", e.Common()
	}
	filename := hash.SHA1(url) + ending
	if files.WriteFile(path.Join(baseFolder, filename), body) != nil {
		return filename, errors.New(files.WriteFile(path.Join(baseFolder, filename), body).Error)
	}
	return filename, nil
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
		newFilm := film.(model.FilmNew)
		for _, id := range newFilm.PersonsID {
			data, err := conn.Get(id, configs.Default.PersonTargetName)
			if err != nil {
				continue
			}
			person := data.(model.Person)
			person.FilmsID = append(person.FilmsID, model.ID(i))
			conn.Upsert(person)
		}
	}

	filmImages, err := parse(configs.Default.FilmImageTargetName)
	if err != nil {
		log.Fatal(err)
	}
	for i, filmImage := range filmImages {
		fmt.Printf("inserting %d film image from %d\n", i, len(filmImages))
		film, e := conn.Get(model.ID(i), configs.Default.FilmTargetName)
		if e != nil {
			log.Fatal(e)
		}
		imagePath, err := uploadAndSaveImage(string(filmImage.(model.Image)), configs.Default.FilmsImageUploadPath)
		if err != nil {
			continue
		}
		f := film.(model.Film)
		f.Images = []model.Image{model.Image(imagePath)}
		conn.Upsert(f)
	}

	personImages, err := parse(configs.Default.PersonImageTargetName)
	if err != nil {
		log.Fatal(err)
	}
	for i, personImage := range personImages {
		fmt.Printf("inserting %d person image from %d\n", i, len(personImages))
		person, e := conn.Get(model.ID(i), configs.Default.PersonTargetName)
		if e != nil {
			return
		}
		imagePath, err := uploadAndSaveImage(string(personImage.(model.Image)), configs.Default.PersonsImageUploadPath)
		if err != nil {
			continue
		}
		p := person.(model.Person)
		p.Images = []model.Image{model.Image(imagePath)}
		conn.Upsert(p)
	}
}

func main() {
	configsPath := flag.String("config", "configs/", "directory with configs")
	flag.Parse()

	err := configuration.UnmarshalConfigs(*configsPath)
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
