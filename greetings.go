package greetings

import "fmt"

func Hello(name string) string{
	//return a message with name
	message := fmt.Sprintf("Hello, %v. Welcome to go modules!", name);
	return message
}