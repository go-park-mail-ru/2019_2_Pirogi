package review

import (
	"time"

	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
)

func CreateReview(in models.NewReview) (models.Review, error) {
	return models.Review{
		NewReview: in,
		Date:      time.Now(),
		Likes:     0,
	}, nil
}
