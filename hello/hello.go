package main 
import (
	"fmt"
	"example.com/greetings"
	"log"
)

func main(){
	// setting the properties for predefinded logger
	// setting log entry prefix
	log.SetPrefix("greetings: ")
	// disable printing the time, source file, and line number
	log.SetFlags(0)


	message,err := greetings.Hello("SHivansh")
	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println(message)
}