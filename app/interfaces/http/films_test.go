package http

//
//import (
//	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
//	mock_database "github.com/go-park-mail-ru/2019_2_Pirogi/app/infrastructure/database/mock"
//	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
//	mockdatabase "github.com/go-park-mail-ru/2019_2_Pirogi/internal/infrastructure/database/mock"
//	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/common"
//	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/validation"
//	"github.com/golang/mock/gomock"
//	"github.com/labstack/echo"
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/require"
//	"net/http"
//	"net/http/httptest"
//	"strings"
//	"testing"
//)
//
//const pathConfigs = "../../../configs"
//
//type TestCaseGetHandlerFilm struct {
//	ParseErrorExpected bool
//	FilmID             model.ID
//	ExpectedFilm       model.Film
//	ExpectedGetError   *model.Error
//	ExpectedPersons    []model.Person
//	ExpectedFindRV     bool
//	ExpectedFullFilm   model.FilmFull
//	ExpectedEchoError  *echo.HTTPError
//}
//
//func TestGetHandlerFilm(t *testing.T) {
//	t.Parallel()
//	ctrl := gomock.NewController(t)
//	mockDb := mock_database.NewMockDatabase(ctrl)
//	testCases := []TestCaseGetHandlerFilm{
//		{
//			FilmID: 0,
//			ExpectedFilm: model.Film{
//				ID:        0,
//				PersonsID: []model.ID{1},
//			},
//			ExpectedPersons: []model.Person{{ID: 1, Name: "actor"}},
//			ExpectedFindRV:  true,
//			ExpectedFullFilm: model.FilmFull{
//				ID:      0,
//				Persons: []model.PersonTrunc{{ID: 1, Name: "actor"}},
//			},
//		},
//		{
//			FilmID:            -1,
//			ExpectedFilm:      model.Film{},
//			ExpectedGetError:  &model.Error{Status: http.StatusNotFound, Error: "no film with the id: -1"},
//			ExpectedEchoError: &echo.HTTPError{Code: http.StatusNotFound, Message: "no film with the id: -1"},
//		},
//		{
//			ParseErrorExpected: true,
//			ExpectedEchoError: &echo.HTTPError{Code: http.StatusNotFound,
//				Message: "strconv.Atoi: parsing \"\": invalid syntax"},
//		},
//	}
//
//	for _, testCase := range testCases {
//		e := echo.New()
//		req := httptest.NewRequest(http.MethodGet, "/", nil)
//		w := httptest.NewRecorder()
//		c := e.NewContext(req, w)
//		c.SetPath("/api/films/:film_id/")
//		c.SetParamNames("film_id")
//
//		if !testCase.ParseErrorExpected {
//			mockDb.EXPECT().Get(testCase.FilmID, configs.Default.FilmTargetName).Return(testCase.ExpectedFilm,
//				testCase.ExpectedGetError)
//			if testCase.ExpectedGetError == nil {
//				mockDb.EXPECT().FindPersonsByIDs(testCase.ExpectedFilm.PersonsID).Return(testCase.ExpectedPersons,
//					testCase.ExpectedFindRV)
//			}
//			c.SetParamValues(testCase.FilmID.String())
//		} else {
//			c.SetParamValues("")
//		}
//
//		err := GetHandlerFilm(mockDb)(c)
//		if testCase.ExpectedEchoError != nil {
//			assert.Error(t, err)
//			assert.Equal(t, testCase.ExpectedEchoError, err)
//		} else {
//			assert.NoError(t, err)
//			assert.Equal(t, w.Code, http.StatusOK)
//			jsonBody, _ := testCase.ExpectedFullFilm.MarshalJSON()
//			assert.True(t, strings.Contains(w.Body.String(), string(jsonBody)))
//		}
//	}
//}
//
//type TestCaseGetHandlerFilmCreate struct {
//	JsonRequestBody       string
//	ParseErrorExpected    bool
//	NewFilm               model.FilmNew
//	ExpectedUpsertError   *model.Error
//	ExpectedFilm          model.Film
//	ExpectedFindFilmRV    bool
//	ExpectedPersons       []model.Person
//	ExpectedFindPersonsRV bool
//	ExpectedEchoError     *echo.HTTPError
//}
//
//func TestGetHandlerFilmCreate(t *testing.T) {
//	validation.InitValidator()
//
//	err := common.UnmarshalConfigs(&pathConfigs)
//	require.NoError(t, err)
//	t.Parallel()
//	ctrl := gomock.NewController(t)
//	mockDb := mockdatabase.NewMockDatabase(ctrl)
//	testCases := []TestCaseGetHandlerFilmCreate{
//		{
//			NewFilm: model.FilmNew{
//				Title:       "Best",
//				Description: "Best description",
//				Year:        2008,
//				Countries:   []string{"Germany"},
//				Genres:      []model.Genre{"боевик"},
//			},
//			ExpectedFilm: model.Film{
//				ID:          0,
//				Title:       "Best",
//				Description: "Best description",
//				Year:        "2008",
//				Countries:   []string{"Germany"},
//				Genres:      []model.Genre{"боевик"},
//			},
//			ExpectedFindFilmRV: true,
//		},
//		{
//			NewFilm: model.FilmNew{
//				Title:       "Best",
//				Description: "Best description",
//				Year:        "2008",
//				Countries:   []string{"Germany"},
//			},
//			ParseErrorExpected: true,
//			ExpectedEchoError:  &echo.HTTPError{Code: http.StatusBadRequest, Message: "genres: Missing required field"},
//		},
//	}
//
//	for _, testCase := range testCases {
//		if testCase.JsonRequestBody == "" {
//			jsonBody, _ := testCase.Newmodel.MarshalJSON()
//			testCase.JsonRequestBody = string(jsonBody)
//		}
//		e := echo.New()
//		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(testCase.JsonRequestBody))
//
//		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
//
//		req.Header.Set(echo.HeaderCookie, "_csrf=test;cinsear_session=test")
//		req.Header.Set(configs.Default.CSRFHeader, "_csrf=test")
//		w := httptest.NewRecorder()
//		c := e.NewContext(req, w)
//		c.SetPath("/api/films/")
//
//		if !testCase.ParseErrorExpected {
//			mockDb.EXPECT().Upsert(testCase.NewFilm).Return(testCase.ExpectedUpsertError)
//			if testCase.ExpectedUpsertError == nil {
//				mockDb.EXPECT().FindFilmByTitle(testCase.Newmodel.Title).Return(testCase.ExpectedFilm,
//					testCase.ExpectedFindFilmRV)
//				if testCase.ExpectedFindFilmRV {
//					mockDb.EXPECT().FindPersonsByIDs(testCase.Expectedmodel.PersonsID).Return(testCase.ExpectedPersons,
//						testCase.ExpectedFindPersonsRV)
//				}
//			}
//		}
//
//		err := GetHandlerFilmCreate(mockDb)(c)
//		if testCase.ExpectedEchoError != nil {
//			assert.Error(t, err)
//			assert.Equal(t, testCase.ExpectedEchoError, err)
//		} else {
//			assert.NoError(t, err)
//			assert.Equal(t, w.Code, http.StatusOK)
//			assert.Equal(t, w.Body.String(), "")
//		}
//	}
//}
