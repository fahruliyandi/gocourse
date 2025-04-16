package basic

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	//length
	str := "hello world"
	fmt.Println(len(str))

	//concate
	str1 := "Bapuk"
	str2 := "Banget"
	result := str1 + str2
	fmt.Println(result)

	//get index
	str3 := "halo"
	fmt.Println(str3[0])

	//substring
	str4 := "fahruliyandi"
	result_str4 := str4[2:5]
	fmt.Println(result_str4)

	//convert int ke string
	num := 100
	str5 := strconv.Itoa(num)
	fmt.Println(str5)

	//string jadi array
	str6 := "apple, banana, grape"
	result_str6 := strings.Split(str6, ",")
	fmt.Println(result_str6)

	//string join
	str7 := []string{"bekasi", "jakarta", "depok"}
	result_str7 := strings.Join(str7, ",")
	fmt.Println(result_str7)

	//contains
	fmt.Println(strings.Contains("Hello", "O"))

	//replace
	fmt.Println(strings.Replace("Fahruliyandi", "Fahrul", "Annisa", 1))

	//trim
	fmt.Println(strings.TrimSpace(" Helloo"))

	//lower and upper
	fmt.Println(strings.ToLower("HELLO"))
	fmt.Println(strings.ToUpper("hello"))

	//count huruf dalam string
	fmt.Println(strings.Count("Hello", "l"))

	//cek prefix dan sufix
	fmt.Println(strings.HasPrefix("Hello", "He"))
	fmt.Println(strings.HasSuffix("Hello", "lo"))

	//regex
	//masukin semua angka dalam string jadi slice
	str8 := "Hel1lo, 123, Go, 11!"
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(str8, -1)
	fmt.Println(matches)

	//STRING BUILDER
	var builder strings.Builder

	//write somestring
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("World")

	result_builder := builder.String()
	fmt.Println(result_builder)

	//reset the builder
	builder.Reset()

}
