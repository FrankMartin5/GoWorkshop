package main

/*
Built-in Data Types
- string
- unsigned integers mean only positives numbers
- int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64

- floating po    int values: float32, float64
- bool
- Boolean operators:  ==, !=, <, >, <=, >=, &&, ||, !

- we can work with pointers also

Visibility
- if its camelCase, its private
- if its TitleCase, its public and exported

- Variables and lambda functions can be:
- Module Scoped
- Function Scooped
- Block Scoped

Numbers
Strings
Collections
- Arrays: fixed length [5]int
- Slices: similar to a dynamic length array, but they are actually chunks of arrays []int
- Maps: key/value dictionaries map[keyType]valueType
- Generics from 1.18
*/

// Global variable
// var url = "https://frontendmasters.com"

var name = "Frontend Masters"
// init can be used to initial things when the app starts

// func init() {
// 	fmt.Println("A")
// }
// func init() {
// 	fmt.Println("B")
// }

func main() {
	// Function Scoped Variables
	// fmt.Println(data.MaxSpeed) // MaxSpeed has to be Titles case so it can be access
	// PrintData()

	print(Area(5, 6))

	// := sets the var for us. Types are implicit so we do not have to declare them.
	// message := "Hello from Go"
	// price := 34.4

	// print(message, price)
	// print("Hello from a module")
	
	// Function can be called because it is listed under the same package
	// Does not need to be imported
}