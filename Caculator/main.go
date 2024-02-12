package main

import "fmt"


func main() {
	var operations string

var number1, number2 int


	fmt.Println("CACULATOR GO 1.0")
	fmt.Println("================")
	fmt.Println("Which operation do you want to perform? (add, subtract, multiply, divide)")
	// If expecting a function to change a variable always use &
	fmt.Scanf("%s", &operations)
	fmt.Println("Enter first number")
	fmt.Scanf("%d", &number1)
	fmt.Println("Enter second number")
	fmt.Scanf("%d", &number2)
	switch operations {
		case "add":
			fmt.Println(number1 + number2)
		case "subtract":
	 		fmt.Println(number1 - number2)
	 	case "multiply":
			fmt.Println(number1 * number2)
		case "divide":
	 		fmt.Println(number1 / number2)
	 }
}



