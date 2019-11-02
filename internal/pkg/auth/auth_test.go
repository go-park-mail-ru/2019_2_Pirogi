package auth

import (
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/common"

	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/user"
	"github.com/stretchr/testify/require"
)

func TestGenerateCookie(t *testing.T) {
	const cookieName = "test"
	const value = "testValue"

	generatedCookie := GenerateCookie(cookieName, value)
	expectedCookie := http.Cookie{
		Name:     cookieName,
		Value:    user.GetMD5Hash(value),
		Expires:  generatedCookie.Expires,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}
	require.Equal(t, generatedCookie, expectedCookie)
}

func TestExpireCookie(t *testing.T) {
	configsPath := "../../../configs"
	err := common.UnmarshalConfigs(&configsPath)
	if err != nil {
		log.Fatal(err.Error())
	}
	const cookieName = "test"
	const value = "testValue"

	generatedCookie := GenerateCookie(cookieName, value)
	require.True(t, generatedCookie.Expires.After(time.Now()))
	ExpireCookie(&generatedCookie)
	require.True(t, generatedCookie.Expires.Before(time.Now()))
}

func TestLogin(t *testing.T) {
}
