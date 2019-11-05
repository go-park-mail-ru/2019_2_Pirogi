package security

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/internal/pkg/models"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestXSSFilterStrings(t *testing.T) {
	input := []string{"<script>alert('you have been pwned')</script>", "<script>console.log('he-he-he')</script>"}
	expected := []string{"&lt;script&gt;alert(&#39;you have been pwned&#39;)&lt;/script&gt;", "&lt;script&gt;console.log(&#39;he-he-he&#39;)&lt;/script&gt;"}
	actual := XSSFilterStrings(input)
	require.Equal(t, expected, actual)
}

func TestXSSFilterGenres(t *testing.T) {
	input := []models.Genre{"<script>alert('you have been pwned')</script>", "<script>console.log('he-he-he')</script>"}
	expected := []models.Genre{"&lt;script&gt;alert(&#39;you have been pwned&#39;)&lt;/script&gt;", "&lt;script&gt;console.log(&#39;he-he-he&#39;)&lt;/script&gt;"}
	actual := XSSFilterGenres(input)
	require.Equal(t, expected, actual)
}

func TestXSSFilterRoles(t *testing.T) {
	input := []models.Role{"<script>alert('you have been pwned')</script>", "<script>console.log('he-he-he')</script>"}
	expected := []models.Role{"&lt;script&gt;alert(&#39;you have been pwned&#39;)&lt;/script&gt;", "&lt;script&gt;console.log(&#39;he-he-he&#39;)&lt;/script&gt;"}
	actual := XSSFilterRoles(input)
	require.Equal(t, expected, actual)
}

func TestCheckNoCSRF(t *testing.T) {

}
