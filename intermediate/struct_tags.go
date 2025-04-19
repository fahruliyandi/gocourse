package intermediate

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName   string `json:"first_name" db:"firstN" xml:"first"`
	LastName    string `json:"last_name"`
	Age         int    `json:"age"`
	IgnoreField string `json:"-"`
}

func main() {

	person := Person{
		FirstName:   "Jhon",
		LastName:    "",
		Age:         30,
		IgnoreField: "Ignore",
	}

	jsonData, err := json.Marshal(person)
	if err != nil {
		log.Fatalln("Error marshalling json:", err)
		return
	}

	fmt.Println(string(jsonData))

}
