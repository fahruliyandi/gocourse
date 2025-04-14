package basic

import "fmt"

func main() {
	message := "hello go"
	message1 := "hello,\tgo" //escape tab
	message2 := "hello,\rgo"
	rawMessage := `Hello\nGo` //`` mengabaikan escape character
	fmt.Println(message)
	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(rawMessage)

	//string adalah array dari character
	message3 := "fahruliyandi"
	fmt.Println(len(message3))

	//char bisa diakses dengan index
	fmt.Println(message3[0]) //return ASCII dari f

	//concate
	greeting := "hello"
	name := "fahrul"
	fmt.Println(greeting + name)

	//compate string
	str1 := "apple"
	str2 := "banana"

	fmt.Println(str1 < str2) //dicompare berdasarkan urutan

	//string iteration seperti iteration pada array atau slice
	for i, char := range message3 {
		fmt.Printf("index %d is %c\n", i, char)
	}

	//string are imutable
	message4 := "hello"
	fmt.Println(&message4)
	message4 = message4 + "fahrul"
	fmt.Println(&message4)
	fmt.Println(message4)

	//rune
	//rune adalah int value merepresentasikan unicode
	var ch rune = 'a'            //single quote
	fmt.Println(ch)              // print ascii number
	fmt.Printf("%c\n", ch)       // print valuenya
	converted_rune := string(ch) //convert rune ke string
	fmt.Println(converted_rune)
	fmt.Printf("type of converted_rune is %T\n", converted_rune) // cek type

}
