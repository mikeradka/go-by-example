// for is Go’s only looping construct. Here are some basic types of for loops.
package main

import "fmt"

func main() {
	fmt.Println("For Loop #1:")
	i := 1 // The most basic type of for loop, with a single condition.
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	fmt.Println("For Loop #2:")
	for j := 0; j < 3; j++ { // A classic initial/condition/after for loop.
		fmt.Println(j)
	}

	fmt.Println("For Loop #3:")
	for i := range 3 { // Another way of accomplishing the basic “do this N times” iteration is range over an integer.
		fmt.Println("range", i)
	}

	fmt.Println("For Loop #4:")
	for {
		fmt.Println("loop")
		break
	}

	fmt.Println("For Loop #5:")
	for n := range 6 { // You can also continue to the next iteration of the loop.
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
