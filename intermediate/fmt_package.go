package main

import "fmt"

func main() {
	fmt.Println("hello")
	fmt.Println(123.45)

	//formating function
	s := fmt.Sprint("hello", " world", 123, 456)
	fmt.Println(s)

}
