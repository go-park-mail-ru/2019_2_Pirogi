package domains

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/validation"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPrepareModelNewUser(t *testing.T) {
	validation.InitValidator()
	userNew := UserNew{}
	userNew.Create("i@artbakulev.com", "1234567890", "Artyom")
	body, err := userNew.MarshalJSON()
	require.NoError(t, err)
	model := UserNew{}
	err = model.Make(body)
	require.NoError(t, err)
	require.Equal(t, userNew, model)
}
