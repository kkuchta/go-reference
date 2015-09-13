// go run reference.go
package main

// Import multiple modules
import (
	"fmt"
	"math/rand"
)

// Constants
const Pi = 3.14

func main() {

	// variables
	first := "hello"
	var second = "world"

	var (
		foo = "foo"
		bar = "bar"
	)
	// Assign to _ to prevent "unused variable" errors.
	_ = foo + bar

	// Printing
	fmt.Println(first, second)
	
	// Random numbers - TODO: how to get a random seed?
	rand.Seed(3)
	fmt.Println("My favorite number is", rand.Intn(10))

	fmt.Println("3 + 4 = ", add(3,4))
	
	sum, diff := sum_and_difference(3, 4)
	fmt.Println("sum, diff = ", sum, diff)

	three, four := add_three_and_four(2)
	fmt.Println("2 + 3, 2 + 4 = ", three, four)

	fmt.Println("pi = ", Pi)
}

// Basic function
func add(x int, y int) int {
	return x + y
}

// Multiple returns
func sum_and_difference(x int, y int) (int, int) {
	return x + y, x - y
}

// Named returns

func add_three_and_four(i int) (three int, four int) {
	three = i + 3
	four = i + 4
	return
}
