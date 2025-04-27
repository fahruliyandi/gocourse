package advance

import "fmt"

func producer(ch chan<- int) {
	for i := range 5 {
		ch <- i
	}
	close(ch)
}

func filter(in <-chan int, out chan<- int) {
	for val := range in {
		if val%2 == 0 {
			out <- val
		}
	}
	close(out)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go producer(ch1)
	go filter(ch1, ch2)

	for val := range ch2 {
		fmt.Println("received ", val)
	}
}

//received from close channel
// func main() {
// 	ch := make(chan int)

// 	close(ch)

// 	val, ok := <-ch
// 	if !ok {
// 		fmt.Println("channel is closed")
// 		return
// 	}
// 	fmt.Println("Receive", val)

// }

//simple closing channel
// func main() {

// 	ch := make(chan int)

// 	go func() {
// 		for i := range 5 {
// 			ch <- i
// 		}
// 		close(ch)
// 	}()

// 	for val := range ch {
// 		fmt.Println("Received ", val)
// 	}

// }
