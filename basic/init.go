package basic

import "fmt"

func init() {
	fmt.Println("Initializing package 1")
}

func init() {
	fmt.Println("Initializing package 2")
}

func init() {
	fmt.Println("Initializing package 3")
}

func main() {
	//init digunakan untuk initialization task sebelum package digunakan
	//sebelum main function digunakan
	fmt.Println("Inside the main function")
}
