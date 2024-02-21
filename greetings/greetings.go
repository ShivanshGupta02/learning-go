package greetings
import (
	"fmt"
	"errors"
	"math/rand"
)
func Hello(name string) (string,error){
	// if no name is given then return error
	if name==""{
		return "", errors.New("empty name")
	}

	// otherwise return the the greeting message
	message := fmt.Sprintf(randomFormat(), name)
	return message,nil
}

func randomFormat() string {
	// a slice of different message formats
	formats := []string{
		"hi, %v. welcome",
		"Nice to onboard you!! %v",
		"Love to see you here!! %v",
	}

	return formats[rand.Intn(len(formats))]
}