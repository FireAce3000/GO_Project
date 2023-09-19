package main
// comment
import (
	"fmt"
	"math"
	"math/rand"
	"math/cmplx"
)
// create variables
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// comment
func main() {
	fmt.Println("A Tour of Go")

	// normal declaration
	var i, j int = 1, 2
	// declaration without var and with :=
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
	// Output: 1 2 3 true false no!

	// math.Pi ausgeben
	fmt.Println(math.Pi)
	// Output: 3.141592653589793

	// math func random Number but max 50
	fmt.Println("My favorite number is", rand.Intn(50))
	// Output: 0 - 49

	// function add
	fmt.Println(add(42, 13))
	// Output: 55

	// function swap
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	// Output: world hello

	// function split
	fmt.Println(split(17))
	// Output: 7 10

	// fmt.Printf - Output with values
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	// Output: Type: bool Value: false
	// Output:Type: uint64 Value: 18446744073709551615
	// Output:Type: complex128 Value: (2+3i)
}

// create function to example
func add(x int, y int) int { // int after the brackets, is the return value
	// return value
	return x + y
}

// create function to switch the words
func swap(x, y string) (string, string) {
	return y, x
}

// create function to return x and y
func split(sum int) (x, y int) { 	// func split (sum int) (int, int) {}
	x = sum * 4 / 9					// var x = sum * 4 / 9
	y = sum - x						// var y = sum - x
	return							// return x, y
}									// }
