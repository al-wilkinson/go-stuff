// Numbers and types
package main

import "fmt"

func main() {

	var x int
	var y int

	x = 1
	y = 2

	fmt.Printf("x = %v, Type: %T\n", x, x)
	fmt.Printf("y = %v, Type: %T\n", y, y)

	mean := (x + y) / 2.0

	fmt.Printf("Mean: %v, Type %T\n", mean, mean)
}
