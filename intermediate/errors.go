package basic

import (
	"errors"
	"fmt"
)

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Math : square foot of negative number")
	}

	return 1, nil
}

func process(data []byte) error {
	if len(data) == 0 {
		return errors.New("Empty data")
	}
	return nil
}

func main() {

	result, err := sqrt(-1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	data := []byte{}
	if errs := process(data); errs != nil {
		fmt.Println(errs)
	} else {
		fmt.Println("Succesfully processed")
	}

	err1 := eprocess()
	if err1 != nil {
		fmt.Println(err1)
	}

	if err2 := readData(); err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("Sukses")
	}

}

// customer errors
type myError struct {
	message string
}

func (e *myError) Error() string {
	return fmt.Sprintf("Error : %s", e.message)
}

func eprocess() error {
	return &myError{
		message: "Custome error message",
	}
}

func readData() error {
	if err := readConfig(); err != nil {
		return fmt.Errorf("Read data : %w", err)
	}

	return nil

}

func readConfig() error {
	return errors.New("Config Error")
}
