package advance

import (
	"fmt"
	"sync"
)

var once sync.Once

func initialize() {
	fmt.Println("Ini cuma di print sekali meskipun bantak go routine yang call function ini")
}

func main() {

	var wg sync.WaitGroup

	for i := range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("Task %d\n", i)
			once.Do(initialize)
		}()
	}

	wg.Wait()
}
