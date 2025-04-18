package basic

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("Helloo \"I am great\"")

	//compile ra regex patter to match email address
	rg := regexp.MustCompile(`[a-zA-Z0-9,_+%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z{2,}]`)

	email1 := "user@email.com"
	email2 := "invalid_email"
	//match
	fmt.Println("Email 1 :", rg.MatchString(email1))
	fmt.Println("Email 2 :", rg.MatchString(email2))

	//capturing groups
	//compile regex to capture date components
	rg = regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)
	date := "2024-07-30"
	submatches := rg.FindStringSubmatch(date)
	fmt.Println(submatches[0])
	fmt.Println(submatches[1])
	fmt.Println(submatches[2])
	fmt.Println(submatches[3])

	//target string
	str := "hello world"
	rg = regexp.MustCompile(`[aiueo]`)
	result := rg.ReplaceAllString(str, "*")
	fmt.Println(result)

	//i - case sensitive
	//m - multi line model
	//s - dot matches all

	rg = regexp.MustCompile(`(?i)go`)
	text := "Golang is great"
	fmt.Println("Match :", rg.MatchString(text))

}
