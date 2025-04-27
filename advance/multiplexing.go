package advance

import (
	"fmt"
	"time"
)

// func main() {
// 	ch1 := make(chan int)
// 	ch2 := make(chan int)

// 	go func() {
// 		ch1 <- 1
// 	}()

// 	go func() {
// 		ch2 <- 2
// 	}()

// 	time.Sleep(time.Millisecond)
// 	for range 2 {
// 		select {
// 		case msg := <-ch1:
// 			fmt.Println("Receive from ch1:", msg)
// 		case msg := <-ch2:
// 			fmt.Println("Receive from ch2:", msg)
// 		default:
// 			fmt.Println("No Channels Ready")
// 		}
// 	}

// }

// =========time out channel==========//
// func main() {
// 	ch := make(chan int)

// 	go func() {
// 		// time.Sleep(2 * time.Second)
// 		ch <- 1
// 		close(ch)
// 	}()

// 	select {
// 	case msg := <-ch:
// 		fmt.Println("Receive from ch :", msg)
// 	case <-time.After(1 * time.Second):
// 		fmt.Println("time out...")
// 	}
// }

// =======close channel pada infinite loop=======
func main() {
	ch := make(chan int, 2)

	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}()

	time.Sleep(time.Second)
	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				fmt.Println("Channel closed")
				return
			}
			fmt.Println("Receive from ch:", msg)
		default:
			fmt.Println("No Channel is ready")
			return
		}
	}
}
