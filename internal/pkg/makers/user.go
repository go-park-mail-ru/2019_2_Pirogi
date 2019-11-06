package makers

import "github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"

func MakeUserTrunc(in models.User) models.UserTrunc {
	return models.UserTrunc{
		ID:          in.ID,
		Username:    in.Username,
		Mark:        in.Mark,
		Description: in.Description,
		Image:       in.Image,
	}
}
