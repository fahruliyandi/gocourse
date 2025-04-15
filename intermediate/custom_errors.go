package basic

import (
	"errors"
	"fmt"
)

func main() {

	err := doSomething()
	if err != nil {
		fmt.Println(err)
	}

	err2 := anotherDoSomething()
	if err2 != nil {
		fmt.Println(err2)
	}

}

type customError struct {
	code    int
	message string
	er      error
}

// error return the error message
// implementing Error() method interface
func (e *customError) Error() string {
	return fmt.Sprintf("Error %d: %s, %s", e.code, e.message, e.er)
}

func doSomething() error {
	return &customError{
		code:    500,
		message: "something went wrong",
	}
}

func doSomethingElse() error {
	return errors.New("Internal error")
}

func anotherDoSomething() error {
	err := doSomethingElse()
	if err != nil {
		return &customError{
			message: "Something went wrong",
			code:    500,
			er:      err,
		}
	}
	return nil
}
