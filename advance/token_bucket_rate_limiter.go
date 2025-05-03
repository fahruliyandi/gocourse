package advance

import (
	"fmt"
	"time"
)

type RateLimiter struct {
	token      chan struct{}
	tokenLimit int
	reffilTime time.Duration
}

func newRateLimiter(limit int, reffilTime time.Duration) *RateLimiter {
	rl := &RateLimiter{
		tokenLimit: limit,
		token:      make(chan struct{}, limit),
		reffilTime: reffilTime,
	}

	for range rl.tokenLimit {
		rl.token <- struct{}{}
	}

	go rl.startReffil()
	return rl
}

func (rl *RateLimiter) startReffil() {
	ticker := time.NewTicker(rl.reffilTime)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			select {
			case rl.token <- struct{}{}:
			default:
			}
		}
	}

}

func (rl *RateLimiter) allow() bool {
	select {
	case <-rl.token:
		return true
	default:
		return false
	}
}

func main() {

	rl := newRateLimiter(5, time.Second)

	for range 10 {
		if rl.allow() {
			fmt.Println("Connection Allowed")
		} else {
			fmt.Println("Connection Denied")
		}
		time.Sleep(200 * time.Millisecond)
	}

}
