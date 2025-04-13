package basic

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Defer")
	fmt.Println("starting main function")

	os.Exit(1)

	fmt.Println("Never be executed")

}
