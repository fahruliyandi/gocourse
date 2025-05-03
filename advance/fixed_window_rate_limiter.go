package advance

import (
	"fmt"
	"sync"
	"time"
)

type RateLimiter struct {
	mu        sync.Mutex
	count     int
	limit     int
	window    time.Duration
	resetTime time.Time
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	rl := &RateLimiter{
		limit:  limit,
		window: window,
	}

	return rl
}

func (rl *RateLimiter) allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	if now.After(rl.resetTime) {
		rl.resetTime = now.Add(rl.window)
		rl.count = 0
	}

	if rl.count < rl.limit {
		rl.count++
		return true
	}

	return false
}

func main() {

	rl := NewRateLimiter(3, time.Second)
	for range 10 {
		if rl.allow() {
			fmt.Println("Request Allowed")
		} else {
			fmt.Println("Request Denied")
		}
		time.Sleep(200 * time.Millisecond)
	}

}
