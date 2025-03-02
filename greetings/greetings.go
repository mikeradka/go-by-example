// https://go.dev/doc/tutorial/create-module

package greetings

import "fmt"

// Hello returns a greeting for the named person.
// 'Hello is the function name, 'name' is a parameter of type string, and the function returns a string.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
