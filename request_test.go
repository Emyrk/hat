package hat

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestURLParams(t *testing.T) {
	req, err := http.NewRequest("GET", "http://google.com", nil)
	require.NoError(t, err)

	URLParams(url.Values{
		"q": []string{"sean"},
	})(t, req)

	require.Equal(t, "http://google.com?q=sean", req.URL.String())
}
