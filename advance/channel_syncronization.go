package advance

import (
	"fmt"
	"strconv"
	"time"
)

// func main() {
// 	done := make(chan int)
// 	go func() {
// 		fmt.Println("working...")
// 		time.Sleep(2 * time.Second)
// 		done <- 0
// 	}()

// 	<-done
// 	fmt.Println("finished")
// }

// func main() {
// 	ch := make(chan int)
// 	go func() {
// 		ch <- 10
// 		// time.Sleep(1 * time.Second)
// 		fmt.Println("sent value")
// 	}()
// 	value := <-ch
// 	fmt.Println(value)
// 	time.Sleep(5 * time.Second)

// }

// ===== go routine syncronization and ensuring all goroutines are completed
// func main() {
// 	numGoroutine := 3
// 	ch := make(chan int, 3)

// 	for i := range numGoroutine {
// 		// time.Sleep(1 * time.Second)
// 		go func(id int) {
// 			fmt.Printf("Working on %d goroutine\n", id)
// 			ch <- id //signaling go routine complete
// 		}(i)
// 	}

// 	for range numGoroutine {
// 		<-ch //wait to each go routing is finished and wait for all go routing to signal completion
// 	}

// 	fmt.Println("All goroutines all complete")
// }

// SYNCRONIZING DATA EXCHANGE
func main() {
	ch := make(chan string)
	go func() {
		for i := range 5 {
			ch <- "Hello " + strconv.Itoa(i)
			time.Sleep(100 * time.Microsecond)
		}
		close(ch)
	}()

	for value := range ch {
		fmt.Println("Received Value ", value, ":", time.Now())
	}
}
