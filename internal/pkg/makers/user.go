package makers

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/domains"
)

func MakeUserTrunc(in domains.User) domains.UserTrunc {
	return domains.UserTrunc{
		ID:          in.ID,
		Username:    in.Username,
		Mark:        in.Mark,
		Description: in.Description,
		Image:       in.Image,
	}
}
