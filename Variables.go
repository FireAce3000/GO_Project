package main
// comment
import (
	"fmt"
)
// comment
func main() {
	fmt.Println("Hello World!")

	var student1 string = "John" //type is string
	var student2 = "Jane" //type is inferred
	x := 2 //type is inferred
	test := "asd"

	fmt.Print(test)
  
	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)

	// Multiple Variable Declaration
	var a, b, c, d int = 1, 3, 5, 7

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	var e, f = 6, "Hello"
  	g, h := 7, "World!"

  	fmt.Println(e)
  	fmt.Println(f)
  	fmt.Println(g)
  	fmt.Println(h)

	var (
		i int
		j int = 1
		k string = "hello"
	)
   
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)

	// Output with Print Function
	var aa,ab string = "Hello","World"

	fmt.Print(aa)
	fmt.Print(ab)

	fmt.Print(aa, "\n")
	fmt.Print(ab, "\n")

	fmt.Print(aa, "\n",ab)

	fmt.Println(aa,ab)

	var ac,ad = 10,20

	fmt.Print(ac,ad)

	var ae string = "Hello"
	var af int = 15
  
	fmt.Printf("ae has value: %v and type: %T\n", ae, ae)
	fmt.Printf("af has value: %v and type: %T", af, af)

	// DataTypes
	var ba bool = true     // Boolean
	var bb int = 5         // Integer
	var bc float32 = 3.14  // Floating point number
	var bd string = "Hi!"  // String
  
	fmt.Println("Boolean: ", ba)
	fmt.Println("Integer: ", bb)
	fmt.Println("Float:   ", bc)
	fmt.Println("String:  ", bd)
}
