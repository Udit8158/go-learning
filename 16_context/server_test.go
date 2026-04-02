package context_server

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// this is like a spy store (for testing)
type StubStore struct {
	response  string
	cancelled bool
}

func (s *StubStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)
	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Println("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

func (s *StubStore) Cancel() {
	s.cancelled = true
}

func (s *StubStore) assertWasCanceled(t *testing.T) {
	t.Helper()
	if !s.cancelled {
		t.Error("store was not told to cancel\n")
	}
}
func (s *StubStore) assertWasNotCanceled(t *testing.T) {
	t.Helper()
	if s.cancelled {
		t.Error("store told to cancel\n")
	}
}

func TestServer(t *testing.T) {
	t.Run("server is giving response", func(t *testing.T) {

		data := "hello, world"
		store := &StubStore{data, false}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		store.assertWasNotCanceled(t)

		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"\n`, response.Body.String(), data)
		}

	})

	t.Run("should tell the store to cancel the work if request is cancelled", func(t *testing.T) {
		data := "Hi there"
		store := &StubStore{data, false}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		// if request got cancelled
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		store.assertWasCanceled(t)
	})
}
