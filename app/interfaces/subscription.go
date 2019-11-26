package interfaces

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/infrastructure/database"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"go.uber.org/zap"
)

type subscriptionRepository struct {
	conn database.Database
}

func (s *subscriptionRepository) Insert(subscriptionNew model.SubscriptionNew) *model.Error {
	return s.conn.Upsert(subscriptionNew)
}

func (s *subscriptionRepository) Update(subscription model.Subscription) *model.Error {
	return s.conn.Upsert(subscription)
}

func (s *subscriptionRepository) Find(userID model.ID) (model.Subscription, *model.Error) {
	raw, err := s.conn.Get(userID, configs.Default.SubscriptionTargetName)
	if err != nil {
		return model.Subscription{}, err
	}
	if subscription, ok := raw.(model.Subscription); !ok {
		return model.Subscription{}, model.NewError(500, "can not cast type")
	} else {
		return subscription, nil
	}

}

func (s *subscriptionRepository) FindSubscriptions(personID model.ID) ([]model.Subscription, *model.Error) {
	return s.conn.FindSubscriptionsOnPerson(personID)
}

func (s *subscriptionRepository) AddEvent(userID model.ID, event model.SubscriptionEvent) *model.Error {
	subscription, err := s.Find(userID)
	zap.S().Debug(subscription)
	if err != nil {
		return err
	}
	subscription.SubscriptionEvents = append(subscription.SubscriptionEvents, event)
	err = s.Update(subscription)
	if err != nil {
		return err
	}
	return nil
}

func (s *subscriptionRepository) SendEventToSubscribers(event model.SubscriptionEvent) *model.Error {
	subscriptions, err := s.FindSubscriptions(event.PersonID)
	zap.S().Debug(subscriptions)
	if err == nil {
		for _, subscription := range subscriptions {
			err = s.AddEvent(subscription.UserID, event)
			if err != nil {
				zap.S().Error(err)
			}
		}
	}
	return err
}

func NewSubscriptionRepository(conn database.Database) *subscriptionRepository {
	return &subscriptionRepository{
		conn: conn,
	}
}
