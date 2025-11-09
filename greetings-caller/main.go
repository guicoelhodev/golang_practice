package main

import(
	"fmt"
	"example/greetings"
)

func main(){
	message := greetings.Hello("Guilherme")
	
	fmt.Println(message)
}

