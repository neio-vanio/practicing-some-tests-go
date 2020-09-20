package request

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRequestWithSelect(t *testing.T) {
	t.Run("fast server", func(t *testing.T) {
		serverFast := newServerTest(time.Second * 0)
		serverSlow := newServerTest(time.Second * 2)

		defer serverFast.Close()
		defer serverSlow.Close()

		urlFast := serverFast.URL
		urlSlow := serverSlow.URL

		expected := urlFast
		result, _ := requestWithSelect(urlSlow, urlFast, time.Second*10)

		if result != expected {
			t.Errorf("result '%s', expected '%s'", result, expected)
		}
	})

	t.Run("timeout server", func(t *testing.T) {
		serverFast := newServerTest(time.Second * 2)
		serverSlow := newServerTest(time.Second * 3)

		defer serverFast.Close()
		defer serverSlow.Close()

		urlFast := serverFast.URL
		urlSlow := serverSlow.URL

		_, err := requestWithSelect(urlSlow, urlFast, time.Second*1)

		if err == nil {
			t.Error("we expected an error")
		}
	})
}

func newServerTest(timeWait time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(timeWait)
		w.WriteHeader(http.StatusOK)
	}))
}
