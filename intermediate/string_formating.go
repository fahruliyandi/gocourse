package basic

import "fmt"

func main() {
	num := 42
	fmt.Printf("%05d\n", num) //00042

	message := "Hello"
	fmt.Printf("|%10s|\n", message)  //|     Hello|
	fmt.Printf("|%-10s|\n", message) //|Hello     |

	//string interpolation
	message2 := "hello \nworld"
	message3 := `hello \nworld`
	fmt.Println(message2)
	fmt.Println(message3)

}
