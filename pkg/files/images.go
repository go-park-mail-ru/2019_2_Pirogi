package files

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/domains"
	"mime"
	"net/http"
)

const MaxUploadSize = 2 * 1024 * 1024

func DetectContentType(data []byte) (ending string, err *domains.Error) {
	fileType := http.DetectContentType(data)
	switch fileType {
	case "image/jpeg", "image/jpg":
	case "image/gif", "image/png":
	case "application/pdf":
		break
	default:
		return "", domains.NewError(http.StatusBadRequest, "unsupported type of file")
	}
	endings, e := mime.ExtensionsByType(fileType)
	if e != nil {
		return "", domains.NewError(http.StatusBadRequest, "can not define extension")
	}
	return endings[0], nil
}




