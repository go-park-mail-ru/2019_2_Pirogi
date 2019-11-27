package model

import "time"

type SubscriptionRequest struct {
	PersonID ID `json:"person_id"`
}

type SubscriptionEvent struct {
	PersonID    ID     `json:"person_id"`
	FilmID      ID     `json:"film_id"`
	Description string `json:"description"`
	Datetime    string `json:"date"`
	IsRead      bool   `json:"is_read"`
}

type SubscriptionEventFull struct {
	Persons     []PersonTrunc `json:"persons"`
	Description string        `json:"description"`
	Datetime    string        `json:"date"`
	IsRead      bool          `json:"is_read"`
}

func NewSubscriptionEvent(personID ID, filmID ID, description string) SubscriptionEvent {
	return SubscriptionEvent{
		FilmID:      filmID,
		PersonID:    personID,
		Description: description,
		Datetime:    time.Now().Format("2006-01-02"),
		IsRead:      false,
	}
}

// Дублирую Subscription, чтобы база понимала, когда обновлять, а когда создавать
type SubscriptionNew struct {
	UserID             ID                  `json:"user_id" bson:"_id"`
	PersonsID          []ID                `json:"person_id"`
	SubscriptionEvents []SubscriptionEvent `json:"subscription_events"`
}

func (sn *SubscriptionNew) String() string {
	return "New subscription of user " + sn.UserID.String()
}

func (sn *SubscriptionNew) ToSubscription() Subscription {
	return Subscription{
		UserID:             sn.UserID,
		PersonsID:          sn.PersonsID,
		SubscriptionEvents: sn.SubscriptionEvents,
	}
}

type Subscription struct {
	UserID             ID                  `json:"user_id" bson:"_id"`
	PersonsID          []ID                `json:"person_id"`
	SubscriptionEvents []SubscriptionEvent `json:"subscription_events"`
}

func NewSubscription(userID ID, personID ID) SubscriptionNew {
	return SubscriptionNew{
		UserID:             userID,
		PersonsID:          []ID{personID},
		SubscriptionEvents: []SubscriptionEvent{},
	}
}

func (s *Subscription) String() string {
	return "Subscription of user " + s.UserID.String()
}

func (se *SubscriptionEvent) String() string {
	return "Subscription event: " + se.Description
}
