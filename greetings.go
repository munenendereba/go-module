package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)
func Hello(name string) (string, error) {
	if (name==""){
		return "",errors.New("Name is blank!")
	}
	//return a message with name
	message := fmt.Sprintf(randomFormat(), name);
	return message,nil
}

func randomFormat() string{
	formats :=[] string{
		"Hi, %v. Welcome!","Hello, %v. How was your day?","%v is this you?"}
	
		return formats[rand.Intn((len(formats)))]
}