// https://go.dev/doc/tutorial/create-module

package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
// 'Hello is the function name, 'name' is a parameter of type string, and the function returns a string.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// If a name was received, return a value that embeds the name
	// in a greeting message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
