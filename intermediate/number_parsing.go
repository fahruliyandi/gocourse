package basic

import (
	"fmt"
	"strconv"
)

func main() {
	numStr := "12345"
	ret, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("error parsing", err)
		return
	}
	fmt.Println("parsed integer :", ret+10)

	ret2, err2 := strconv.ParseInt(numStr, 10, 32)
	if err2 != nil {
		fmt.Println("error parsing", err2)
		return
	}

	fmt.Println("parsed integer :", ret2+10)

	floatStr := "3.14"
	floatVal, err3 := strconv.ParseFloat(floatStr, 64)
	if err3 != nil {
		fmt.Println("error parsing", err3)
		return
	}

	fmt.Printf("Parsed float: %.2f\n", floatVal)

}
