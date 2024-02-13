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

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
    // A map to associate names with messages.
    messages := make(map[string]string)
    // Loop through the received slice of names, calling
    // the Hello function to get a message for each name.
    for _, name := range names {
        message, err := Hello(name)
        if err != nil {
            return nil, err
        }
        // In the map, associate the retrieved message with
        // the name.
        messages[name] = message
    }
    return messages, nil
}