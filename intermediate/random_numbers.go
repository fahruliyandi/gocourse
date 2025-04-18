package basic

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	// rn := rand.New(rand.NewSource(42))
	// fmt.Println(rand.Intn(5) + 1) // generate random number dari 1 - 5

	//random floating number
	// fmt.Println(rand.Float64()) //default between 0.0 and 1.0

	for {
		fmt.Println("Welcome to dice game")
		fmt.Println("1.Roll the dice")
		fmt.Println("2.Exit")
		fmt.Println("Enter your choice (1 or 2) :")

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil || (choice != 1 && choice != 2) {
			fmt.Println("Invalid choice, please enter 1 or 2 ")
			break
		}

		if choice == 2 {
			fmt.Println("Thank and good bye...")
			break
		}

		die1 := rand.Intn((6) + 1)
		die2 := rand.Intn((6) + 1)

		fmt.Printf("You rolled %d and %d", die1, die2)
		fmt.Println("Total is", die1+die2)

		fmt.Println("Do you want to roll again? [y/n]")
		var rollAgain string
		_, err1 := fmt.Scan(&rollAgain)
		if err1 != nil || (strings.ToLower(rollAgain) != "y" && strings.ToLower(rollAgain) != "n") {
			println("Invalid input, Assuming no", err1)
			break
		}

		if strings.ToLower(rollAgain) == "n" {
			fmt.Println("Thanks for playing, good bye...")
			break
		}

	}

	// var nomor int
	// nomor = 10
	// fmt.Println(nomor)
	// var nomor2 *int
	// nomor2 = &nomor
	// *nomor2 = 100
	// fmt.Println(nomor)
	// fmt.Println(nomor2)

}
