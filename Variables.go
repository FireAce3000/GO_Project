package main
// comment
import (
	"fmt"
	"math" // math package
	"math/rand" // random
	"math/cmplx" // complex math (sqrt)
	"time" // timepackage
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

	// time 
	var today = time.Now()
	fmt.Println(today)
	// Output: 2023-09-19 12:32:01.470252403 +0000 UTC m=+0.000031059

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

	// for loop
	sum := 1						// var erstellen
	for i := sum; i < 1000; i++ {		// for erstellen with components by ;
		sum += sum					// what is to do
		}
	fmt.Println(sum)
	// Output: 1024

	// while = for
	sum1 := 1
	for sum1 < 1000 {
		sum1 += sum1
	}
	fmt.Println(sum1)
	// Output: 1024

	// if LIKE for
	x := 5
	if x > 0 {
		fmt.Println("x is bigger than 0")
	} else {
		fmt.Println("x is smaller than 0")
	}
	// Output: x is bigger than 0

	// switch
	var join int = 1
	switch join {
	case 1: fmt.Print("case1")
	case 2: fmt.Print("case2")
	default: fmt.Print("casedefault")
	}
	// Output: case1

	// switch without condition - switch true
	t := time.Now() 
	switch { 								// empty condition
	case t.Hour() < 12:						// example 
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	// defer - "führt den Befehl erst aus, wenn der Befehl schon mal ausgeführt wurde"
	defer fmt.Println("world")
	fmt.Println("hello")
	// Output: hello world

	// defer stack
	for i := 0; i < 10; i++ {
		defer fmt.Print(i)
	}
	fmt.Print("done ")
	// Output: done 9 8 7 6 5 4 3 2 1 0
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
