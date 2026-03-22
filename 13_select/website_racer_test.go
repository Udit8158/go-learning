package websiteracer

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestWebsiteRacer(t *testing.T) {
	t.Run("test the faster url", func(t *testing.T) {
		slowMockServer := createMockServerWithDealy(1) // 1s delay
		fastMockServer := createMockServerWithDealy(0) // 0s delay

		defer slowMockServer.Close()
		defer fastMockServer.Close()

		fmt.Println("url", fastMockServer.URL)
		got, _ := WebsiteRacer(slowMockServer.URL, fastMockServer.URL)

		if got != fastMockServer.URL {
			t.Errorf("Expected %q but got %q", fastMockServer.URL, got)
		}

	})
	t.Run("test for the error after 10 sec delay", func(t *testing.T) {
		slowMockServer := createMockServerWithDealy(11) //11s delay
		fastMockServer := createMockServerWithDealy(12) //10s delay

		defer slowMockServer.Close()
		defer fastMockServer.Close()

		_, err := WebsiteRacer(slowMockServer.URL, fastMockServer.URL)

		if err == nil {
			t.Errorf("Expected error but got no error\n")
		}
	})
}

func createMockServerWithDealy(delay int64) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Duration(delay) * time.Second)
		w.WriteHeader(http.StatusOK)
	}))
}
