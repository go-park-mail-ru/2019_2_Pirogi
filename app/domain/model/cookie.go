package model

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/hash"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

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
	c.Cookie.Value = hash.SHA1(value)
	c.Cookie.HttpOnly = true
	c.Cookie.Path = "/"
	c.Cookie.SameSite = http.SameSiteStrictMode
}
