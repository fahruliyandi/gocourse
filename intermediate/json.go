package intermediate

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name    string  `json:"name"`
	Age     int     `json:"age,omitempty"`
	Email   string  `json:"email"`
	Address Address `json:"address"`
}

type Address struct {
	City   string `json:"city"`
	Nation string `json:"nation"`
}

func main() {
	person := Person{
		Name:  "Jhon",
		Age:   30,
		Email: "123@email.com",
	}

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling to json", err)
		return
	}
	fmt.Println(string(jsonData))

	person1 := Person{
		Name:  "Recoba",
		Age:   40,
		Email: "recoba@haha.com",
		Address: Address{
			City:   "New York",
			Nation: "Uruguay",
		},
	}

	jsonData1, err := json.Marshal(person1)
	if err != nil {
		fmt.Println("Error encoding person1 :", err)
		return
	}
	fmt.Println(string(jsonData1))
	jsonData2 := `{
		"name" : "vieri",
		"empId" : "00098",
		"age" : 40,
		"address" : {
			"city" : "milan",
			"nation" : "italia"
		} 
	}`

	var employeeFromJson Employee
	err = json.Unmarshal([]byte(jsonData2), &employeeFromJson)
	if err != nil {
		fmt.Println("Error unmarshalling json", err)
	}

	fmt.Println(employeeFromJson.Name)
	fmt.Println(employeeFromJson.Age)
	fmt.Println(employeeFromJson.EmpId)
	fmt.Println(employeeFromJson.Address.City)
	fmt.Println(employeeFromJson.Address.Nation)

	listOfCity := []Address{
		{City: "London", Nation: "UK"},
		{City: "Berlin", Nation: "German"},
	}

	fmt.Println(listOfCity)
	jsonList, err := json.Marshal(listOfCity)
	if err != nil {
		log.Fatalln("Error marshalling list of city:", err)
		return
	}

	fmt.Println(string(jsonList))

	// handling unkown json structure
	jsonData3 := `{"name" :"Ronaldo", "age" : 30, "address" : {"city":"Bogota","nation":"Colombia"}}`
	var jsonMap map[string]interface{}
	err = json.Unmarshal([]byte(jsonData3), &jsonMap)
	if err != nil {
		fmt.Println("Error unmarshaling unknown json structure :", err)
		return
	}

	fmt.Println(jsonMap)

	fmt.Println(jsonMap["address"].(map[string]interface{})["city"])
	fmt.Println(jsonMap["name"])
	fmt.Println(jsonMap["age"])

	intMap := map[string]map[string]int{
		"atas": {
			"bawah": 10,
		},
	}

	fmt.Println(intMap["atas"]["bawah"])

}

type Employee struct {
	Name    string  `json:"name"`
	EmpId   string  `json:"empId"`
	Age     int     `json:"age"`
	Address Address `json:"address"`
}
