package basic

import (
	"fmt"
)

type Person struct {
	firtName string
	lastName string
	age      int
	address  Address
	PhoneCell
}

type Address struct {
	city    string
	country string
}

type PhoneCell struct {
	cellNumber string
}

func (p Person) fullName() string {
	return p.firtName + " " + p.lastName
}

func (p *Person) changeFirstName(name string) {
	p.firtName = name
}

func main() {

	p := Person{
		firtName: "fahrul",
		lastName: "liyandi",
		age:      35,
		address: Address{
			city:    "London",
			country: "Uganda",
		},
		PhoneCell: PhoneCell{
			cellNumber: "000111",
		},
	}

	fmt.Println(p.fullName())

	p.changeFirstName("bapuk")
	fmt.Println(p.fullName())
	fmt.Println(p.address.city)
	fmt.Println(p.address.country)
	fmt.Println(p.cellNumber)

	p1 := Person{
		firtName: "Ajib",
		age:      37,
	}

	p1.address.city = "Birmingham"
	p1.address.country = "Kenya"

	fmt.Println(p.firtName)
	fmt.Println(p1.age)

	//annonimous struct
	user := struct {
		username string
		email    string
	}{
		username: "fahruliayndi",
		email:    "123@gmail.com",
	}

	fmt.Println(user)

}
