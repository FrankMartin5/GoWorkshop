package main

import "fmt"

/* Packages must be in the same folder
- If the function name is upper case it is public
- If its lower then its private
*/
func PrintData()  {
	
	fmt.Print("Hello")
	fmt.Println("World")
	fmt.Println(name)
}


func Area(num1 int, num2 int) int {
	return num1 * num2
}

