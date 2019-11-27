package interfaces

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/infrastructure/database"
)

type listsRepository struct {
	conn database.Database
}

func (l listsRepository) Insert(listNew model.ListNew) *model.Error {
	return l.conn.Upsert(listNew)
}

func (l listsRepository) Update(list model.List) *model.Error {
	return l.conn.Upsert(list)
}

func (l listsRepository) Get(id model.ID) (model.List, *model.Error) {
	list, ok := l.conn.FindListByID(id)
	if !ok {
		return model.List{}, model.NewError(404, "Список не найден")
	}
	return list, nil
}

func (l listsRepository) GetByUserID(userID model.ID) ([]model.List, *model.Error) {
	lists, err := l.conn.FindListsByUserID(userID)
	if err != err {
		return nil, err
	}
	return lists, nil
}

func NewListsRepository(conn database.Database) *listsRepository {
	return &listsRepository{
		conn: conn,
	}
}
