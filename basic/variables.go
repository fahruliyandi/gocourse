package basic

import "fmt"

var middle_name = "middle" //global variable

func main() {
	// var age int
	// var name = "fahrul"
	// var name1 = "liyandi"

	// count := 10
	// fullName := "fahruliyandi"
	middle_name = "changed"
	fmt.Println(middle_name)
	//Default values
	//Numeric type = 0
	//Boolean = false
	//String = ""
	//Pointers,slice,maps,functions and structs:nil

}

func print_middle_name() {
	fmt.Println(middle_name)
}
