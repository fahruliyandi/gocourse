package advance

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type atomicCounter struct {
	count int64
}

func (ac *atomicCounter) increment() {
	atomic.AddInt64(&ac.count, 1)
}

func (ac *atomicCounter) getValue() int64 {
	return atomic.LoadInt64(&ac.count)
}

func main() {
	var wg sync.WaitGroup
	numGoroutine := 10
	ac := &atomicCounter{}

	wg.Add(numGoroutine)
	// value := 0

	for range numGoroutine {
		go func() {
			defer wg.Done()
			for range 1000 {
				ac.increment()
				// value++
			}
		}()
	}

	wg.Wait()

	fmt.Println("Final result :", ac.getValue())
	// fmt.Println("Final result :", value)

}
