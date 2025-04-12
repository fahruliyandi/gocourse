package basic

import "fmt"

type employee struct {
	FirtName string
	LastName string
	Age      int
}

func main() {
	//PascalCase
	//untuk composite data type
	//Structs, interface, enums

	//snake_case
	//untuk variable
	//user_id, first_name, http_request

	//UPPERCASE
	//untuk naming constants
	const MAXRETRIES = 5

	//camelCase
	//untuk nama variable juga bisa

	var employee_id = 1001
	fmt.Println(employee_id)

}
