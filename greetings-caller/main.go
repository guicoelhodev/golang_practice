package main

import (
	"fmt"
	"time"
	//"example/greetings"
	//"log"
)

func main(){
	loop()
}

func loop() {
	start := time.Now()

	items := make([]int, 0, 1000000)

	for i := 0; i < 1000000; i++ {
		items = append(items, i)
	}

	elapsed := time.Since(start)

	fmt.Println("Time taken:", elapsed)  // Time taken: 1.827185ms
}

//func greetingFn(){
//	message, err := greetings.Hello("Guilherme")
//
//	log.SetPrefix("greetings: ")
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println(message)
//}
