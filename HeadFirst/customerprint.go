package main

import "fmt"

func main() {
	quantity := 4
	// quantity = 4 if we forget :, it will be treated as an assignment, not a declaration and we cant assignto a variable.
	// quantity := 4 It will fail because u cna declare a variable only once
	// quantity := "4" //Static typed
	length, width := 1.2, 2.4
	// length, width := 1.2 // must provide value for every variable
	customerName := "Damon Cole"

	// all declatred variables must be used in our program. if any variable is unused, declared variables must be omitted
	fmt.Println(customerName)
	fmt.Println("has ordered", quantity, "sheets")
	fmt.Println("each with an area of")
	fmt.Println(length*width, "square meters")
}
