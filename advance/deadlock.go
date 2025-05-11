package advance

import (
	"fmt"
	"sync"
)

func main() {
	var m1, m2 sync.Mutex

	go func() {
		m1.Lock()
		fmt.Println("Go Routine 1 m1 locked")
		m2.Lock()
		fmt.Println("Go Routine 1 m2 locked")
		m1.Unlock()
		m2.Unlock()
	}()

	go func() {
		m2.Lock()
		fmt.Println("Go Routine 2 m2 locked")
		m1.Lock()
		fmt.Println("Go Routine 2 m1 locked")
		m2.Unlock()
		m1.Unlock()
	}()

	select {}
}
