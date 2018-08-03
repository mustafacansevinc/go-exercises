package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

func main() {
	go say("Hello World")
	// the keyword go means run as a seperate task

	// so the rest of the program runs
	fmt.Println("I need some sleep")
	time.Sleep(100 * time.Millisecond)
}