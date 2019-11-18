package domains

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/user"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type CookieRepository interface {
	Insert(cookie Cookie) (ID, error)
	Update(cookie Cookie) error
	Delete(id ID) bool
	Get(id ID) Cookie
	GetFromRequest(r *http.Request, name string) (*Cookie, error)
	SetOnResponse(res *echo.Response, r *Cookie)
	Find(cookie *Cookie) ID
}

type Cookie struct {
	UserID ID           `json:"user-id" bson:"_id" valid:"numeric"`
	Cookie *http.Cookie `valid:"cookie"`
}

func (c *Cookie) String() string {
	return c.Cookie.Value
}

func (c *Cookie) Expire() {
	c.Cookie.Expires = time.Unix(0, 0)
	c.Cookie.Path = "/"
}

func (c *Cookie) Generate(cookieName, value string) {
	expiration := time.Now().Add(configs.Default.CookieAuthDurationHours * time.Hour)
	c.Cookie.Name = cookieName
	c.Cookie.Expires = expiration
	c.Cookie.Value = user.Hash(value)
	c.Cookie.HttpOnly = true
	c.Cookie.Path = "/"
	c.Cookie.SameSite = http.SameSiteStrictMode
}
