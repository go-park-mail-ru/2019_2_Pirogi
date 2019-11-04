package validators

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidateEmailOK(t *testing.T) {
	const email = "artbakulev@gmail.com"
	require.True(t, ValidateEmail(email))
}

func TestValidateEmailFail(t *testing.T) {
	const email = "1@om"
	require.False(t, ValidateEmail(email))
}
