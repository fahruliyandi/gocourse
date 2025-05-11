package advance

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// func printNumbers() {
// 	for i := range 5 {
// 		fmt.Println(time.Now())
// 		fmt.Println(i)
// 		time.Sleep(500 * time.Millisecond)
// 	}
// }

// func printLetters() {

// 	for _, letter := range "ABCDE" {
// 		fmt.Println(time.Now())
// 		fmt.Println(string(letter))
// 		time.Sleep(500 * time.Millisecond)
// 	}

// }

func heavyTask(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Printf("Task %d start processing\n", id)
	for range 100_000_000 {

	}
	fmt.Println(time.Now())
	fmt.Printf("Task %d is finished\n", id)

}

func main() {

	// go printNumbers()
	// go printLetters()

	// time.Sleep(3 * time.Second)

	var wg sync.WaitGroup
	numThreads := 4

	runtime.GOMAXPROCS(numThreads)

	for i := range numThreads {
		wg.Add(1)
		go heavyTask(i, &wg)
	}

	wg.Wait()
}
