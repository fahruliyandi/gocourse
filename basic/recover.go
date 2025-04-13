package basic

import "fmt"

func main() {

	process()
	fmt.Println("Return from process")

}

func process() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("recovered : ", r)
		}
	}()

	fmt.Println("start process")
	panic("Something went wrong")
}
