package advance

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)
	go func() {
		ch <- 1
		fmt.Println("Go Routine 2 Second")
		time.Sleep(2 * time.Second)
	}()

	go func() {
		fmt.Println("Go Routine 3 Second")
		ch <- 2
		time.Sleep(3 * time.Second)
	}()

	receiver := <-ch
	fmt.Println(receiver)
	fmt.Println("End Of Program")

}
