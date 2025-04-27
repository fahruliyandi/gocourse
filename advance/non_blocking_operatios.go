package advance

import (
	"fmt"
	"time"
)

func main() {

	// ch := make(chan int)

	// //NON BLOCKING RECEIVE FROM CHANNEL
	// select {
	// case msg := <-ch:
	// 	fmt.Println("Received ", msg)
	// default:
	// 	fmt.Println("Channel is not ready to send")
	// }

	// //NON BLOCKING SEND TO CHANNEL
	// select {
	// case ch <- 1:
	// 	fmt.Println("Sukses send to channel")
	// default:
	// 	fmt.Println("channel is not ready to receive")
	// }

	data := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case d := <-data:
				fmt.Println("Received ", d)
			case <-quit:
				fmt.Println("Stoping...")
				return
			default:
				fmt.Println("waiting for data...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	for i := range 5 {
		data <- i
		time.Sleep(time.Second)
	}

	quit <- true

}
