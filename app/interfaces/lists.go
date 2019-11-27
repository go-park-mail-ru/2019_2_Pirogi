package interfaces

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/infrastructure/database"
)

type listsRepository struct {
	conn database.Database
}

func (l listsRepository) Insert(list model.List) *model.Error {
	return l.conn.Upsert(list)
}

func (l listsRepository) Update(list model.List) *model.Error {
	panic("implement me")
}

func (l listsRepository) Get(id model.ID) (model.List, *model.Error) {
	panic("implement me")
}

func (l listsRepository) GetByUserID(userID model.ID) (model.List, *model.Error) {
	panic("implement me")
}

func NewListsRepository(conn database.Database) *listsRepository {
	return &listsRepository{
		conn: conn,
	}
}
