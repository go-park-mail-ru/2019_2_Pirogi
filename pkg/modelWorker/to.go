package modelWorker

import "github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"

func GenresToStrings(genres []model.Genre) (result []string) {
	for _, genre := range genres {
		result = append(result, string(genre))
	}
	return result
}
