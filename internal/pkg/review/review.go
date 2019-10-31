package review

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"time"
)

func CreateReview(in models.NewReview) (models.Review, error) {
	return models.Review{
		ID:        -1,
		NewReview: in,
		Date:      time.Now(),
		Likes:     0,
	}, nil
}
