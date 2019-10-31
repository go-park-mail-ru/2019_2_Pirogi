package validators

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateEmailOK(t *testing.T) {
	const email = "artbakulev@gmail.com"
	require.True(t, ValidateEmail(email))
}

func TestValidateEmailFail(t *testing.T) {
	const email = "1@om"
	require.False(t, ValidateEmail(email))
}
