/*
	Display the following menu
		1. Add
		2. Subtract
		3. Multiply
		4. Divide
		5. Exit
	Accept the user choice
	If the user choice is 1 - 4
		accept 2 numbers from the user
		perform the corresponding operation
		print the result
		display the menu again

	If the user choice = 5
		print "Thank you"
		exit the application

	If the user choice is NOT 1 - 5
		print "Invalid choice"
		display the menu again
*/

package main

import "fmt"

func main() {
	var n1, n2, result, userChoice int
	for {
		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Println("5. Exit")
		fmt.Println("Enter your choice:")
		fmt.Scanln(&userChoice)

		if userChoice == 5 {
			fmt.Println("Thank you!")
			break
		}
		if userChoice < 1 || userChoice > 5 {
			fmt.Println("Invalid Choice")
			continue
		}

		fmt.Println("Enter the operands :")
		fmt.Scanln(&n1, &n2)
		switch userChoice {
		case 1:
			result = n1 + n2
		case 2:
			result = n1 - n2
		case 3:
			result = n1 * n2
		case 4:
			result = n1 / n2
		}
		fmt.Println("Result :", result)
	}
}
