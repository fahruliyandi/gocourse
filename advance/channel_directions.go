package advance

import "fmt"

func main() {

	ch := make(chan int)

	producer(ch)
	consumer(ch)
}

//send only channel
func producer(ch chan<- int) {
	go func() {
		for i := range 5 {
			ch <- i
		}
		close(ch)
	}()
}

//receive only channel
func consumer(ch <-chan int) {
	for value := range ch {
		fmt.Println("Value :", value)
	}
}
