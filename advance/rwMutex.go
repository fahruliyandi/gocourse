package advance

import (
	"fmt"
	"sync"
	"time"
)

var counter int
var rwmu sync.RWMutex

func readCounter(wg *sync.WaitGroup) {
	defer wg.Done()
	defer rwmu.RUnlock()
	rwmu.RLock()
	fmt.Printf("Counter : %d\n", counter)
}

func writeCounter(wg *sync.WaitGroup, value int) {
	defer wg.Done()
	defer rwmu.Unlock()
	rwmu.Lock()
	counter = value
	fmt.Printf("Written value %d for counter\n", value)
}

func main() {

	var wg sync.WaitGroup
	for range 5 {
		wg.Add(1)
		go readCounter(&wg)
	}

	wg.Add(1)
	time.Sleep(time.Second)
	go writeCounter(&wg, 18)

	wg.Wait()

}
