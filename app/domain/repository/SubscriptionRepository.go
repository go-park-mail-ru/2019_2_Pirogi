package repository

import "github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"

type SubscriptionRepository interface {
	Insert(subscription model.SubscriptionNew) *model.Error
	Update(subscription model.Subscription) *model.Error
	Find(userID model.ID) (model.Subscription, *model.Error)
	FindSubscriptions(personID model.ID) ([]model.Subscription, *model.Error)
	AddEvent(userID model.ID, event model.SubscriptionEvent) *model.Error
	SendEventToSubscribers(event model.SubscriptionEvent) *model.Error
}
