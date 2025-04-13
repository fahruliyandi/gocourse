package basic

import "fmt"

func main() {
	//panic(interface{})

	// process(10)
	process(-1)

}

func process(input int) {

	defer fmt.Println("Defer 1")
	defer fmt.Println("Defer 2")

	if input < 0 {
		fmt.Println("Before Panic")
		panic("Cannot be negative number")
	}
	fmt.Printf("Processed number is %d", input)
}
