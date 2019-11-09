package common

import (
	"bufio"
	"bytes"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/validators"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestReadLinesInvalidFile(t *testing.T) {
	var invalidImage []byte
	_, err := WriteFileWithGeneratedName(invalidImage, "")
	require.Error(t, err)
}

func TestReadLinesInvalidPath(t *testing.T) {
	const testFileName = "./test.png"
	const invalidPath = "./@@@@"
	file, err := os.Open(testFileName)
	defer func() { _ = file.Close() }()
	require.NoError(t, err)
	validImage, err := ioutil.ReadAll(file)
	require.NoError(t, err)
	_, err = WriteFileWithGeneratedName(validImage, invalidPath)
	require.Error(t, err)
}

func TestReadLines(t *testing.T) {
	const testFileName = "./test.png"
	const validPath = "./"
	file, err := os.Open(testFileName)
	defer func() { _ = file.Close() }()
	require.NoError(t, err)
	validImage, err := ioutil.ReadAll(file)
	require.NoError(t, err)
	name, err := WriteFileWithGeneratedName(validImage, validPath)
	require.NoError(t, err)
	require.NotNil(t, name)
	err = os.Remove(name)
	require.NoError(t, err)
}

func TestMakeJSONArray(t *testing.T) {
	input := [][]byte{{'a', 'b', 'c'}, {'d', 'e', 'f'}}
	expected := []byte{'[', 'a', 'b', 'c', ',', 'd', 'e', 'f', ']'}
	actual := MakeJSONArray(input)
	require.Equal(t, expected, actual)
}

func TestNormalizePath(t *testing.T) {
	actual := "/api/users"
	expected := "/api/users/"
	NormalizePath(&actual)
	require.Equal(t, expected, actual)
}

func TestReadBody(t *testing.T) {
	input := []byte{'a', 'b', 'c', 'd'}
	expected := []byte{'a', 'b', 'c', 'd'}
	e := echo.New()
	reader := bufio.NewReader(bytes.NewBuffer(input))
	req := httptest.NewRequest("POST", "http://www.google.com", reader)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	actual, err := ReadBody(ctx)
	require.Nil(t, err)
	require.Equal(t, expected, actual)
}

func TestPrepareModelNewPerson(t *testing.T) {
	validators.InitValidator()
	userNew := models.NewPerson{
		Name:       "Artyom",
		Roles:      []models.Role{"actor", "director"},
		Birthday:   "09.12.1998",
		Birthplace: "USA",
	}
	body, err := userNew.MarshalJSON()
	require.NoError(t, err)
	model, err := PrepareModel(body, userNew)
	require.NoError(t, err)
	require.Equal(t, userNew, model)
}

func TestPrepareModelNewFilm(t *testing.T) {
	validators.InitValidator()
	filmNew := models.NewFilm{
		Title:       "Matrix",
		Description: "Lalasdadasdasdl",
		Year:        "1998",
		Countries:   []string{"USA"},
		Genres:      []models.Genre{"драма"},
		PersonsID:   []models.ID{},
	}
	body, err := filmNew.MarshalJSON()
	require.NoError(t, err)
	model, err := PrepareModel(body, filmNew)
	require.NoError(t, err)
	require.Equal(t, filmNew, model)
}

func TestPrepareModelNewReview(t *testing.T) {
	validators.InitValidator()
	reviewNew := models.NewReview{
		Title:    "Matrix",
		Body:     "asdasdasdasd",
		FilmID:   2,
		AuthorID: 3,
	}
	body, err := reviewNew.MarshalJSON()
	require.NoError(t, err)
	model, err := PrepareModel(body, reviewNew)
	require.NoError(t, err)
	require.Equal(t, reviewNew, model)
}

func TestPrepareModelNewUser(t *testing.T) {
	validators.InitValidator()
	userNew := models.NewUser{
		Credentials: models.Credentials{
			Email:    "i@artbakulev.com",
			Password: "1234567890",
		},
		Username: "Artyom",
	}
	body, err := userNew.MarshalJSON()
	require.NoError(t, err)
	model, err := PrepareModel(body, userNew)
	require.NoError(t, err)
	require.Equal(t, userNew, model)
}

func TestCheckPOSTRequest(t *testing.T) {
	err := UnmarshalConfigs("../../../configs")
	require.NoError(t, err)
	e := echo.New()
	cookieCSRF := &http.Cookie{
		Name:  "_csrf",
		Value: "test",
		Path:  "/",
	}
	cookieAuth := &http.Cookie{
		Name:  "cinsear_session",
		Value: "test",
	}
	reader := bufio.NewReader(nil)
	req := httptest.NewRequest("POST", "http://cinsear.ru", reader)
	req.AddCookie(cookieCSRF)
	req.AddCookie(cookieAuth)
	req.Header.Set(configs.Default.CSRFHeader, "_csrf=test")
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	session, err := CheckPOSTRequest(ctx)
	require.NoError(t, err)
	require.Equal(t, "test", session.Value)
}

func TestUnmarshalConfigs(t *testing.T) {
	err := UnmarshalConfigs("../../../configs")
	require.NoError(t, err)
	require.Equal(t, "_csrf", configs.Default.CSRFCookieName)
}
