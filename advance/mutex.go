package advance

import (
	"fmt"
	"sync"
)

// =========mutex bukan pada struct=============
func main() {
	var counter int
	var wg sync.WaitGroup
	var mx sync.Mutex
	numGoroutines := 5

	increment := func() {
		defer wg.Done()
		for range 1000 {
			mx.Lock()
			counter++
			mx.Unlock()
		}

	}

	for range numGoroutines {
		wg.Add(1)
		go increment()
	}

	wg.Wait()

	fmt.Printf("Result %d\n", counter)
}

// ==========mutex pada struct=============
// type counter struct {
// 	mx    sync.Mutex
// 	count int
// }

// func (c *counter) increment() {
// 	c.mx.Lock()
// 	defer c.mx.Unlock()
// 	c.count++
// }

// func (c *counter) getCount() int {
// 	c.mx.Lock()
// 	defer c.mx.Unlock()
// 	return c.count
// }

// func main() {

// 	var wg sync.WaitGroup
// 	counts := counter{}
// 	numWorkers := 10

// 	wg.Add(numWorkers)

// 	for range numWorkers {
// 		go func(wg *sync.WaitGroup) {
// 			defer wg.Done()
// 			for range 1000 {
// 				counts.increment()
// 				// counts.count++
// 			}
// 		}(&wg)
// 	}

// 	wg.Wait()
// 	fmt.Println("Count : ", counts.getCount())

// }
