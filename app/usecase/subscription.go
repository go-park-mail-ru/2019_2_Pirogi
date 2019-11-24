package usecase

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/repository"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/json"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/modelWorker"
	"github.com/labstack/echo"
)

type SubscriptionUsecase interface {
	GetUserByContext(ctx echo.Context) (model.User, *model.Error)
	GetPersonsTruncListJSONBlob(userID model.ID) ([]byte, *model.Error)
	GetNewEventsListJSONBlob(userID model.ID) ([]byte, *model.Error)
	Unsubscribe(userID, personID model.ID) *model.Error
	Subscribe(userID, personID model.ID) *model.Error
}

func NewSubscriptionUsecase(subscriptionRepository repository.SubscriptionRepository,
	cookieRepository repository.CookieRepository, personRepository repository.PersonRepository,
	userRepository repository.UserRepository) *subscriptionUsecase {
	return &subscriptionUsecase{
		subscriptionRepo: subscriptionRepository,
		cookieRepo:       cookieRepository,
		personRepo:       personRepository,
		userRepo:         userRepository,
	}
}

type subscriptionUsecase struct {
	subscriptionRepo repository.SubscriptionRepository
	cookieRepo       repository.CookieRepository
	personRepo       repository.PersonRepository
	userRepo         repository.UserRepository
}

func (u *subscriptionUsecase) GetUserByContext(ctx echo.Context) (model.User, *model.Error) {
	return u.cookieRepo.GetUserByContext(ctx)
}

func (u *subscriptionUsecase) GetPersonsTruncListJSONBlob(userID model.ID) ([]byte, *model.Error) {
	subscription, err := u.subscriptionRepo.Find(userID)
	if err != nil {
		return nil, err
	}
	persons := u.personRepo.GetMany(subscription.PersonsID)
	personsTrunc := modelWorker.TruncPersons(persons)
	body := modelWorker.MarshalPersonsTrunc(personsTrunc)
	jsonBlob := json.MakeJSONArray(body)
	return jsonBlob, nil
}

func (u *subscriptionUsecase) GetNewEventsListJSONBlob(userID model.ID) ([]byte, *model.Error) {
	subscription, err := u.subscriptionRepo.Find(userID)
	if err != nil {
		return nil, err
	}
	var newSubscriptionEvents []model.SubscriptionEvent
	for _, event := range subscription.SubscriptionEvents {
		if !event.IsRead {
			newSubscriptionEvents = append(newSubscriptionEvents, event)
		}
	}
	body := modelWorker.MarshalSubscriptionEvents(newSubscriptionEvents)
	jsonBlob := json.MakeJSONArray(body)
	return jsonBlob, nil
}

func (u *subscriptionUsecase) Unsubscribe(userID, personID model.ID) *model.Error {
	subscription, err := u.subscriptionRepo.Find(userID)
	if err != nil {
		return err
	}
	var isDeleted bool
	for idx, id := range subscription.PersonsID {
		if id == personID {
			subscription.PersonsID = append(subscription.PersonsID[:idx], subscription.PersonsID[idx+1:]...)
			isDeleted = true
			break
		}
	}
	if !isDeleted {
		return model.NewError(400, "Пользователь не подписан на этого актера")
	}
	err = u.subscriptionRepo.Update(subscription)
	if err != nil {
		return err
	}
	return nil
}

func (u *subscriptionUsecase) Subscribe(userID, personID model.ID) *model.Error {
	var isAlreadySubscribed = false
	subscription, err := u.subscriptionRepo.Find(userID)
	if err != nil && err.Status == 404 {
		subscriptionNew := model.NewSubscription(userID, personID)
		err := u.subscriptionRepo.Insert(subscriptionNew)
		if err != nil {
			return err
		}
		subscription = subscriptionNew.ToSubscription()
		isAlreadySubscribed = true
	}

	if err != nil && err.Status != 404 {
		return err
	}

	for _, id := range subscription.PersonsID {
		if id == personID {
			isAlreadySubscribed = !isAlreadySubscribed
			break
		}
	}
	if isAlreadySubscribed {
		return model.NewError(400, "Пользователь уже подписан на этого актера")
	}

	subscription.PersonsID = append(subscription.PersonsID, personID)
	err = u.subscriptionRepo.Update(subscription)
	if err != nil {
		return err
	}
	return nil
}
