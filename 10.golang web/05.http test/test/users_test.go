package test

import (
	"httptest/router"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

var url = ""

func TestMain(m *testing.M) {
	server := httptest.NewServer(router.Routers())
	defer server.Close()
	url = server.URL
	m.Run()
}

func TestUsersEndpoint(t *testing.T) {
	res, err := http.Get(url + "/users")
	if err != nil {
		require.Error(t, err)
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	require.Equal(t, "semua data users", string(body))

}
