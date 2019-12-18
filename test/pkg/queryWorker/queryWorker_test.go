package queryWorker

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/configuration"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/queryWorker"
	"github.com/go-park-mail-ru/2019_2_Pirogi/test/fixture"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetPipelineForMongoByContext(t *testing.T) {
	err := configuration.UnmarshalConfigs("../../../configs/")
	require.NoError(t, err)
	ctx := fixture.NewEchoContext(nil,
		map[string]string{"year_min": "1998", "year_max": "2018", "genres": "драма,мелодрама"})
	pipeline := queryWorker.GetPipelineForMongoByContext(ctx, configs.Default.FilmTargetName)
	require.Equal(t, 5, len(pipeline))
}
