package domains

import "net/http"

type CookieRepository interface {
	Insert(cookie Cookie) (ID, error)
	Update(cookie Cookie) error
	Delete(id ID) bool
	Get(id ID) Cookie
}

type Cookie struct {
	UserID ID           `json:"user-id" bson:"_id" valid:"numeric"`
	Cookie *http.Cookie `valid:"cookie"`
}
