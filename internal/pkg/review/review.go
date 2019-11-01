package review

import (
	"time"

	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
)

func CreateReview(id models.ID, in models.NewReview) (models.Review, error) {
	return models.Review{
		NewReview: in,
		ID:        id,
		Date:      time.Now(),
		Likes:     0,
	}, nil
}
