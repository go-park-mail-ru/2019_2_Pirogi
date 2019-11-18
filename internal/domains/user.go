package domains

import (
	"github.com/asaskevich/govalidator"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
)

type UserRepository interface {
	Insert(newUser UserNew) (ID, error)
	Update(id ID, user User) error
	Delete(id ID) bool
	Get(id ID) User
	GetMany(target Target, id ID) []User
	GetByCookie(cookie Cookie) (User, bool)
	GetByEmail(email string) (User, bool)
	MakeTrunc(user User) UserTrunc
}

type UserNew struct {
	Email    string `json:"email" valid:"email"`
	Password string `json:"password" valid:"password"`
	Username string `json:"username" valid:"stringlength(2|50)"`
}

func (un *UserNew) Make(body []byte) error {
	err := un.UnmarshalJSON(body)
	if err != nil {
		return err
	}
	_, err = govalidator.ValidateStruct(un)
	return err
}

type UserTrunc struct {
	ID          ID     `json:"id" valid:"numeric"`
	Username    string `json:"username" valid:"title"`
	Mark        Mark   `json:"mark" valid:"mark, optional"`
	Description string `json:"description" valid:"description"`
	Image       Image  `json:"image" valid:"image, optional"`
}

type User struct {
	ID          ID     `json:"id" valid:"numeric"`
	Email       string `json:"email" valid:"email"`
	Password    string `json:"password" valid:"password"`
	Username    string `json:"username" valid:"title"`
	Mark        Mark   `json:"mark" valid:"mark, optional"`
	Description string `json:"description" valid:"description"`
	Image       Image  `json:"image" valid:"image, optional"`
}

func (u *User) CheckPassword(password string) bool {
	return u.Password == password
}

func (un *UserNew) Create(email, password, username string) {

}

func (un *UserNew) ToUser(id ID) User {
	return User{
		ID:          id,
		Email:       un.Email,
		Password:    un.Password,
		Username:    un.Username,
		Mark:        0,
		Description: "",
		Image:       configs.Default.DefaultImageName,
	}
}

func (u *User) Trunc() UserTrunc {
	return UserTrunc{
		ID:          u.ID,
		Username:    u.Username,
		Mark:        u.Mark,
		Description: u.Description,
		Image:       u.Image,
	}
}
