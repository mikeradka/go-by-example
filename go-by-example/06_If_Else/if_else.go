package main

import "fmt"

func main() {
	//basic example of if/else
	if 7%2 == 0 {
		fmt.Println("7 is event")
	} else {
		fmt.Println("7 is odd")
	}

	// you can have an if statement without an else
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// logical operators like && and || are often useful in conditions
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	// A statement can precede conditionals; any variables
	// declared in this statement are available in the current
	// and all subsequent branches
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has mulitple digits")
	}

	// You don't need parentheses around conditions, but braces
	// are required.
	// There is no ternary if in Go, so you'll need to use a full
	// if statement even for basic conditions.
}
