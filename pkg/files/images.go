package files

import (
	"mime"
	"net/http"

	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
)

const MaxUploadSize = 2 * 1024 * 1024

func DetectContentType(data []byte) (ending string, err *model.Error) {
	fileType := http.DetectContentType(data)
	switch fileType {
	case "image/jpeg", "image/jpg":
	case "image/gif", "image/png":
	case "application/pdf":
		break
	default:
		return "", model.NewError(http.StatusBadRequest, "unsupported type of file")
	}
	endings, e := mime.ExtensionsByType(fileType)
	if e != nil {
		return "", model.NewError(http.StatusBadRequest, "can not define extension")
	}
	return endings[0], nil
}
