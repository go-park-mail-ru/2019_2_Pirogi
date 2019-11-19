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
	"io/ioutil"
	"net/http"
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
	case configs.Default.FilmImageTargetName, configs.Default.PersonImageTargetName:
		var images []models.Image
		for dec.More() {
			var image models.Image
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
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	return common.WriteFileWithGeneratedName(body, baseFolder)
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

	filmImages, err := parse(configs.Default.FilmImageTargetName)
	if err != nil {
		log.Fatal(err)
	}
	for i, filmImage := range filmImages {
		fmt.Printf("inserting %d film image from %d\n", i, len(filmImages))
		film, e := conn.Get(models.ID(i), configs.Default.FilmTargetName)
		if e != nil {
			log.Fatal(e)
		}
		imagePath, err := uploadAndSaveImage(string(filmImage.(models.Image)), configs.Default.FilmsImageUploadPath)
		if err != nil {
			log.Fatal(err)
		}
		f := film.(models.Film)
		f.Images = []models.Image{models.Image(imagePath)}
		conn.Upsert(f)
	}

	personImages, err := parse(configs.Default.PersonImageTargetName)
	if err != nil {
		log.Fatal(err)
	}
	for i, personImage := range personImages {
		fmt.Printf("inserting %d person image from %d\n", i, len(personImages))
		person, e := conn.Get(models.ID(i), configs.Default.PersonTargetName)
		if e != nil {
			continue
		}
		imagePath, err := uploadAndSaveImage(string(personImage.(models.Image)), configs.Default.PersonsImageUploadPath)
		if err != nil {
			continue
		}
		p := person.(models.Person)
		p.Images = []models.Image{models.Image(imagePath)}
		conn.Upsert(p)
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
