package auth

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/user"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"time"
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
	const cookieName = "test"
	const value = "testValue"

	generatedCookie := GenerateCookie(cookieName, value)
	require.True(t, generatedCookie.Expires.After(time.Now()))
	ExpireCookie(&generatedCookie)
	require.True(t, generatedCookie.Expires.Before(time.Now()))
}

func TestLogin(t *testing.T) {
}
