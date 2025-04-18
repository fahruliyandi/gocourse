package basic

import (
	"bufio"
	"fmt"
	"html/template"
	"os"
	"strings"
)

func main() {
	// tmp1 := template.New("example")
	// template.New()
	tmp1, err := template.New("example").Parse("Welcome, {{.name}}! How are you doing?\n")
	if err != nil {
		panic(err)
	}

	//define the data
	data := map[string]interface{}{
		"name": "Jhon",
	}

	data["name"] = "Fahrul"

	err = tmp1.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name :")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	//define name template
	templates := map[string]string{
		"welcome":      "welcome, {{.name}} we are glad you join",
		"notification": "{{.name}} you have a new notification {{.notification}}",
		"error":        "Oops! an error occured: {{.errorMessage}}",
	}

	//parse and store template
	parsedTemplate := make(map[string]*template.Template)

	for key, value := range templates {
		parsedTemplate[key] = template.Must(template.New(key).Parse(value))
	}

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1.Join")
		fmt.Println("2.Get Notification")
		fmt.Println("3.Get Error:")
		fmt.Println("4.Exit:")
		fmt.Println("Choose an option")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		data := make(map[string]interface{})
		var temp1 *template.Template

		switch choice {
		case "1":
			temp1 = parsedTemplate["welcome"]
			// data = map[string]interface{}{"name": name}
			data["name"] = name
		case "2":
			fmt.Println("Enter your notification")
			notif, _ := reader.ReadString('\n')
			temp1 = parsedTemplate["notification"]
			data = map[string]interface{}{
				"name":         name,
				"notification": notif,
			}
		case "3":
			fmt.Println("Enter your error message : ")
			errMsg, _ := reader.ReadString('\n')
			temp1 = parsedTemplate["error"]
			data = map[string]interface{}{
				"name":         name,
				"errorMessage": errMsg,
			}
		case "4":
			println("Exitig....")
			return
		default:
			fmt.Println("Please choose 1, 2, 3 or 4")
			continue

		}

		err := temp1.Execute(os.Stdout, data)
		if err != nil {
			fmt.Println("Error")
		}
	}

}
