package basic

import (
	"fmt"
)

func main() {

	// fmt.Println(sum(1, 2, 3, 4))
	statement, total := sum("The total of 1,2,3,4 is ", 1, 2, 3, 4)
	fmt.Println(statement, total)

	numSlice := []int{1, 2, 3, 4}
	statement1, total1 := sum("total", numSlice...)
	fmt.Println(statement1, total1)

}

func sum(statement string, nums ...int) (string, int) {
	total := 0
	for _, v := range nums {
		total += v
	}

	return statement, total
}
