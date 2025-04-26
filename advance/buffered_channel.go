package advance

import (
	"fmt"
	"time"
)

// func main() {
// 	//======= BLOCKING ON RECEIVE ONLY IF THE BUFFER IS EMPTY =========
// 	ch := make(chan int, 2)
// 	go func() {
// 		time.Sleep(10 * time.Second)
// 		ch <- 1
// 	}()
// 	fmt.Println("Value", <-ch)

// }

func main() {
	// ========= BLOCKING SEND ONLY IF THE BUFFER IS FULL===========
	//make(chan type, capacity)
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println("Receiving from buffer")
	go func() {
		fmt.Println("2 second start")
		time.Sleep(2 * time.Second)
		fmt.Println("received :", <-ch)
		// fmt.Println("blocking ends here")
	}()
	// fmt.Println("start blocking here")
	ch <- 3 //blocking karena pada saat sending value 3 buffer sudah full dan menunggu go routine selesai dimana didalam go routing ada receive channel yang membuat channel tidak full lagi
	// fmt.Println("received :", <-ch)
	// fmt.Println("received :", <-ch)
	// fmt.Printf("Buffered Channel")
}
