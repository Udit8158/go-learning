package counter_sync

import "sync"

type Counter struct {
	count int
	mu    sync.Mutex
}

func (c *Counter) Value() int {
	return c.count
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}
