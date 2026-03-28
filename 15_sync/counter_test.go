package counter_sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, &counter, 3)
	})
	t.Run("increment 3 times concurrently, leaves it at 3", func(t *testing.T) {
		counter := Counter{}
		var wg sync.WaitGroup
		wantedCount := 1000

		for range wantedCount {
			wg.Go(func() {
				counter.Inc()
			})
		}

		wg.Wait()
		assertCounter(t, &counter, 1000)
	})
}

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
