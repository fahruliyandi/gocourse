package basic

import "fmt"

func main() {

	fmt.Println(factorial(5))

}

func factorial(n int) int {
	//base case : factorial 0 is 1
	if n == 0 {
		return 1
	}
	//rescursive case
	return n * factorial(n-1)

}
