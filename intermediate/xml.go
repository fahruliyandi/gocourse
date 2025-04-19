package intermediate

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Person struct {
	XMLName xml.Name `xml:"person"` //untuk root element
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
	City    string   `xml:"city"`
	Email   string   `xml:"email"`
}

type Employee struct {
	XMLName xml.Name        `xml:"employee"`
	Name    string          `xml:"name"`
	Address EmployeeAddress `xml:"address"`
}

type EmployeeAddress struct {
	City    string `xml:"city"`
	Country string `xml:"country"`
}

func main() {

	person := Person{
		Name:  "Jhon",
		Age:   30,
		City:  "London",
		Email: "Email@Email.com",
	}

	xmlData, err := xml.MarshalIndent(person, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling data into xml :", err)
	}

	fmt.Println(string(xmlData))

	xmlRaw := `<person>
					<name>Jeni</name>
					<age>20</age>
					<city>Glasgow</city>
					<email>jemi@Email.com</email>
			  </person>`

	var person1 Person
	err = xml.Unmarshal([]byte(xmlRaw), &person1)
	if err != nil {
		log.Fatalln("Error unmarshaling xml : ", err)
		return
	}

	fmt.Println(person1.Name)
	fmt.Println(person1.Age)
	fmt.Println(person1.Email)
	fmt.Println(person1.City)

	employee := Employee{
		Name: "Eddie",
		Address: EmployeeAddress{
			City:    "San Fransisco",
			Country: "United States",
		},
	}

	empRaw, err := xml.MarshalIndent(employee, "", "  ")
	if err != nil {
		log.Fatalln("Error marshalling employee xml:", err)
	}

	fmt.Println(string(empRaw))

	book := Book{
		Title: "Go Bootcamp",
		ISBN:  "ISBN",
	}

	xmlBook, err := xml.MarshalIndent(book, "", "  ")
	if err != nil {
		log.Fatalln("Error marshalling book to xml", err)
	}

	fmt.Println(string(xmlBook))

}

type Book struct {
	XMLName xml.Name `xml:"name"`
	ISBN    string   `xml:"isbn,attr"`
	Title   string   `xml:"title,attr"`
}
