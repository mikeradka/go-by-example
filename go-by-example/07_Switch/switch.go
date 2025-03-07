package main

import (
	"fmt"
	"time"
)

func main() {
	// A basic switch statement
	i := 2
	fmt.Print("Write ", i, " as ") // note the different uses of Print vs Println
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Print("two")
	case 3:
		fmt.Println("three")
	}
	fmt.Println(".")

	// You can use commas to separate multiple expressions in the same case
	// statement. We use the optional default case in this example as well.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend.")
	default:
		fmt.Println("It's a weekday.")
	}

	// switch without an expression is an alternate way to express if/else
	// logic. Here we also show how the case expressions can
	// be non-constants.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon.")
	default:
		fmt.Println("It's after noon.")
	}

	// A type switch compares types instead of values. You can use
	// this to discover the type of an interface value. In this
	// example, the variable t will have the type corresponding to
	// its clause.
	whatAmI := func(i any) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool.")
		case int:
			fmt.Println("I'm an int.")
		default:
			fmt.Printf("Dont know type %T.\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
