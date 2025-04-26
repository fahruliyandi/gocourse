package intermediate

import (
	"fmt"
	"time"
)

func sayHello() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello from go routine")
}

func printNumber() {
	for i := 1; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i)
	}
}

func printLetters() {
	for _, letter := range "abcd" {
		fmt.Println(string(letter))
		time.Sleep(200 * time.Millisecond)
	}
}

func doWork() error {
	time.Sleep(1 * time.Second)
	return fmt.Errorf("error in do work")
}

func main() {

	var err error

	fmt.Println("Befor hello")
	go sayHello()
	fmt.Println("After Hello")

	go func() {
		err = doWork()
	}()
	go printNumber()
	go printLetters()

	time.Sleep(2 * time.Second)

	if err != nil {
		fmt.Println("Error :", err)
	}

}
