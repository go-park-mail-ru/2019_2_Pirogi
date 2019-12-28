package validation

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/validation"
	"github.com/go-park-mail-ru/2019_2_Pirogi/test/fixture"
	"testing"

	valid "github.com/asaskevich/govalidator"
	"github.com/stretchr/testify/require"
)

func TestValidation(t *testing.T) {
	validation.InitValidator()
	_, err := valid.ValidateStruct(fixture.Rating)
	require.NoError(t, err)

	_, err = valid.ValidateStruct(fixture.UserTrunc)
	require.NoError(t, err)

	_, err = valid.ValidateStruct(fixture.FilmTrunc)
	require.NoError(t, err)

	_, err = valid.ValidateStruct(fixture.UserCredentials)
	require.NoError(t, err)

	_, err = valid.ValidateStruct(fixture.PersonTrunc)
	require.NoError(t, err)

	_, err = valid.ValidateStruct(fixture.Person)
	require.NoError(t, err)

	_, err = valid.ValidateStruct(fixture.Film)
	require.NoError(t, err)

	_, err = valid.ValidateStruct(fixture.Review)
	require.NoError(t, err)
}
