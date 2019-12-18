package fixture

import "github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"

var UserID = model.ID(1)
var Email = "test@test.com"
var Username = "test"
var Description = "description"
var Image = model.Image("cffa50a32cb13a240d705317bcec65dd1f31b6ad.jpg")
var Password = "c07c99b2d15100b3597997b65d0de82d6051ed56"

var User = model.User{
	ID:          UserID,
	Email:       Email,
	Password:    Password,
	Username:    Username,
	Description: Description,
	Image:       Image,
}

var UserNew = model.UserNew{
	Email:    Email,
	Password: Password,
}

var UserTrunc = model.UserTrunc{
	ID:          UserID,
	Username:    Username,
	Description: Description,
	Image:       Image,
}

var UserCredentials = model.UserCredentials{
	Email:    Email,
	Password: Password,
}
