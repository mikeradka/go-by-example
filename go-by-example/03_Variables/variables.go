package main

import "fmt"

func main() {
	var a = "initial" // declare a single variable
	fmt.Println(a)

	var b, c int = 1, 2 // declare multiple vars at once
	fmt.Println(b, c)

	var d = true
	fmt.Println(d) // go will infer initialized variables

	var e int
	fmt.Println(e) // variables declared without a corresponding initialization are zero-valued.

	f := "apple"
	fmt.Println(f) // := is shorthand for declaring and initializing a variable
}
