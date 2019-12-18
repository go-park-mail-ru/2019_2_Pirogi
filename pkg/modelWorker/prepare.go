package modelWorker

import (
	"github.com/asaskevich/govalidator"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/hash"
	"html"
)

func PrepareUserNew(userNew *model.UserNew) *model.Error {
	_, err := govalidator.ValidateStruct(userNew)
	if err != nil {
		return model.NewError(400, err.Error())
	}
	userNew.Password = hash.SHA1(html.EscapeString(userNew.Password))
	userNew.Email = html.EscapeString(userNew.Email)
	return nil
}

func PrepareUserUpdate(userUpdate *model.User) *model.Error {
	_, err := govalidator.ValidateStruct(userUpdate)
	userUpdate.Username = html.EscapeString(userUpdate.Username)
	userUpdate.Email = html.EscapeString(userUpdate.Email)
	userUpdate.Description = html.EscapeString(userUpdate.Description)
	if err != nil {
		return model.NewError(400, err.Error())
	}
	return nil
}


