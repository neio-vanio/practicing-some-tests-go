package context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response  string
	cancelled bool
}

func (s *SpyStore) Fetch() string {
	time.Sleep(1 * time.Second)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func TestHandler(t *testing.T) {
	t.Run("simple test without cancel option", func(t *testing.T) {
		data := "hello, man"
		svr := Server(&SpyStore{response: data})

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("result: '%s', expected: '%s'", response.Body.String(), data)
		}
	})

	t.Run("simple test with cancel option", func(t *testing.T) {
		data := "hello, man"

		spyStore := &SpyStore{response: data}
		svr := Server(spyStore)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Second, cancel)
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("result: '%s', expected: '%s'", response.Body.String(), data)
		}

		if spyStore.cancelled {
			t.Errorf("store could not have been canceled")
		}

	})

}
