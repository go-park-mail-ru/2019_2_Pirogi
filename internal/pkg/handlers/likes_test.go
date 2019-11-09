package handlers

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/common"
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/validators"
	"github.com/stretchr/testify/require"
	"testing"
)

type TestCaseGetHandlerLikesCreate struct {
}

func TestGetHandlerLikesCreate(t *testing.T) {
	validators.InitValidator()
	err := common.UnmarshalConfigs("../../../configs")
	require.NoError(t, err)
	t.Parallel()
}

