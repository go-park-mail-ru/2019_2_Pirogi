package images

import (
	"errors"
	"mime"
	"net/http"
	"os"
	"path/filepath"
)

const MaxUploadSize = 2 * 1024 * 1024

// TODO: change for server
const UploadUsersPath = "/Users/artyombakulev/Projects/2019_2_Pirogi/media/users/"

func DetectContentType(data []byte) (ending string, err error) {
	fileType := http.DetectContentType(data)
	switch fileType {
	case "image/jpeg", "image/jpg":
	case "image/gif", "image/png":
	case "application/pdf":
		break
	default:
		return "", errors.New("unsupported type of file")
	}
	endings, err := mime.ExtensionsByType(fileType)
	return endings[0], nil
}

func GenerateFilename(target, userID, ending string) string {
	return target + "_" + userID + ending
}

func WriteFile(fileBytes []byte, filename, path string) error {
	newPath := filepath.Join(path, filename)
	newFile, err := os.Create(newPath)
	if err != nil {
		return err
	}
	if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
		return err
	}
	return nil
}

func GetFields(r *http.Request) (ID, loadTarget string, err error) {
	ID = r.PostFormValue("id")
	if ID == "" {
		return "", "", errors.New("specify ID")
	}

	loadTarget = r.PostFormValue("target")
	if loadTarget == "" {
		return "", "", errors.New("set the load target")
	}
	return ID, loadTarget, nil
}
