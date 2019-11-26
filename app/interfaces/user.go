package interfaces

import (
	"github.com/asaskevich/govalidator"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/infrastructure/database"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/hash"
	"html"
)

type userRepository struct {
	conn database.Database
}

func (u userRepository) Insert(newUser model.UserNew) *model.Error {
	return u.conn.Upsert(newUser)
}

func (u userRepository) Update(user model.User) *model.Error {
	return u.conn.Upsert(user)
}

func (u userRepository) Delete(id model.ID) *model.Error {
	//TODO: база!!! хочу по айди удалять
	return u.conn.Delete(id)
}

func (u userRepository) Get(id model.ID) (model.User, *model.Error) {
	userInterface, e := u.conn.Get(id, configs.Default.UserTargetName)
	if e != nil {
		return model.User{}, e
	}
	if user, ok := userInterface.(model.User); !ok {
		return model.User{}, model.NewError(500, "can not cast cookie")
	} else {
		return user, nil
	}
}

func (u userRepository) GetMany(ids []model.ID) []model.User {
	return u.conn.FindUsersByIDs(ids)
}

func (u userRepository) GetByEmail(email string) (model.User, *model.Error) {
	return u.conn.FindUserByEmail(email)
}

func (u userRepository) PrepareUserNew(userNew *model.UserNew) *model.Error {
	_, err := govalidator.ValidateStruct(userNew)
	if err != nil {
		return model.NewError(400, err.Error())
	}
	userNew.Password = hash.SHA1(html.EscapeString(userNew.Password))
	userNew.Email = html.EscapeString(userNew.Email)
	return nil
}

func NewUserRepository(conn database.Database) *userRepository {
	return &userRepository{
		conn: conn,
	}
}
