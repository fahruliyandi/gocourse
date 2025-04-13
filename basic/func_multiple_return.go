package basic

import (
	"errors"
	"fmt"
)

func main() {

	q, r := Divide(10, 3)
	fmt.Printf("Quotient : %v, Remainder : %v\n", q, r)

	result, err := Compare(3, 3)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println(result)
	}

}

func Divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

func Compare(a, b int) (string, error) {
	if a > b {
		return "a greater then b", nil
	} else if a < b {
		return "a lower then b", nil
	} else {
		return "", errors.New("Unable to compare which is greater")
	}

}
