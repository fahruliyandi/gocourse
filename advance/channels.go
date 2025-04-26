package advance

import (
	"fmt"
	"time"
)

func main() {

	//variable := make(chan type)
	greeting := make(chan string)
	greetString := "Hello World"

	go func() {
		greeting <- greetString //blocking because it is countinuesly trying to receive data
		greeting <- "ribet amat"
		for _, v := range "abcde" {
			greeting <- "Alphabet " + string(v)
		}
	}()

	// go func() {
	// 	receiver := <-greeting
	// 	fmt.Println(receiver)
	// 	receiver = <-greeting
	// 	fmt.Println(receiver)
	// }()

	receiver := <-greeting
	fmt.Println(receiver)
	receiver = <-greeting
	fmt.Println(receiver)
	for range 5 {
		receiver = <-greeting
		fmt.Println(receiver)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("end of program")
}
