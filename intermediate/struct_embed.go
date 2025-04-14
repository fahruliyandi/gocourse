package basic

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) introduce() {
	fmt.Printf("Halo saya %v dan umur saya %v", p.name, p.age)
}

type Employee struct {
	Person
	empId  int
	salary float64
}

//override method
func (e Employee) introduce() {
	fmt.Printf("Halo saya employee %v dan umur saya %v, gaji saya %v", e.name, e.age, e.salary)

}

func main() {
	emp := Employee{
		Person: Person{
			name: "fahrul",
			age:  35,
		},
		empId:  101010,
		salary: 10000000,
	}

	fmt.Println(emp.name)
	fmt.Println(emp.age)
	fmt.Println(emp.empId)
	fmt.Println(emp.salary)
	emp.introduce()

}
