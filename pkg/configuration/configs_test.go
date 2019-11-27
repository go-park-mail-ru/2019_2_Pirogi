package configuration

import (
	"testing"

	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/stretchr/testify/require"
)

func TestUnmarshalConfigs(t *testing.T) {
	configPath := "../configs"
	err := UnmarshalConfigs(configPath)
	require.NoError(t, err)
	require.Equal(t, "_csrf", configs.Default.CSRFCookieName)
}
