package model

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/hash"
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

func (c *Cookie) CopyFromCommon(cookie *http.Cookie) {
	c.Cookie = cookie
}

func (c *Cookie) GenerateAuthCookie(id ID, cookieName, value string) {
	cookie := &http.Cookie{
		Name:       cookieName,
		Value:      hash.SHA1(value),
		Path:       "/",
		Expires:    time.Now().Add(configs.Default.CookieAuthDurationHours * time.Hour),
		HttpOnly:   true,
		SameSite:   http.SameSiteStrictMode,
	}
	c.Cookie = cookie
	c.UserID = id
}
