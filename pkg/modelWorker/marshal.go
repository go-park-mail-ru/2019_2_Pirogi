package modelWorker

import "github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"

func MarshalFilmsTrunc(filmsTrunc []model.FilmTrunc) (body [][]byte) {
	for _, filmTrunc := range filmsTrunc {
		raw, err := filmTrunc.MarshalJSON()
		if err != nil {
			continue
		}
		body = append(body, raw)
	}
	return body
}

func MarshalPersonsTrunc(personsTrunc []model.PersonTrunc) (body [][]byte) {
	for _, personTrunc := range personsTrunc {
		raw, err := personTrunc.MarshalJSON()
		if err != nil {
			continue
		}
		body = append(body, raw)
	}
	return body
}

func MarshalReviews(reviews []model.Review) (body [][]byte) {
	for _, review := range reviews {
		raw, err := review.MarshalJSON()
		if err != nil {
			continue
		}
		body = append(body, raw)
	}
	return body
}

func MarshalReviewsFull(reviewsFull []model.ReviewFull) (body [][]byte) {
	for _, reviewFull := range reviewsFull {
		raw, err := reviewFull.MarshalJSON()
		if err != nil {
			continue
		}
		body = append(body, raw)
	}
	return body
}

func MarshalTrailers(trailers []model.TrailerWithTitle) (body [][]byte) {
	for _, trailer := range trailers {
		raw, err := trailer.MarshalJSON()
		if err != nil {
			continue
		}
		body = append(body, raw)
	}
	return body
}

func MarshalSubscriptionEvents(events []model.SubscriptionEvent) (body [][]byte) {
	for _, event := range events {
		raw, err := event.MarshalJSON()
		if err != nil {
			continue
		}
		body = append(body, raw)
	}
	return body
}
