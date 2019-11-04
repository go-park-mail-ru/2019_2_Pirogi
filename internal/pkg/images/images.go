package images

import (
	"mime"
	"net/http"

	Error "github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/error"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/user"
)

const MaxUploadSize = 2 * 1024 * 1024

func DetectContentType(data []byte) (ending string, err *models.Error) {
	fileType := http.DetectContentType(data)
	switch fileType {
	case "image/jpeg", "image/jpg":
	case "image/gif", "image/png":
	case "application/pdf":
		break
	default:
		return "", Error.New(http.StatusBadRequest, "unsupported type of file")
	}
	endings, e := mime.ExtensionsByType(fileType)
	if e != nil {
		return "", Error.New(http.StatusBadRequest, "can not define extension")
	}
	return endings[0], nil
}

func GenerateFilename(salt, userID, ending string) string {
	return user.GetMD5Hash(salt+userID) + ending
}
