package main

import(
	"fmt"
	"log"
	"example/greetings"
)

func main(){
	message, err := greetings.Hello("Guilherme")

	log.SetPrefix("greetings: ")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}

