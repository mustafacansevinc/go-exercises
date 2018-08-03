package main

import (
	"fmt"
)

func say(s string, c chan int) {
	var i int
	for i := 0; i < 5; i++ {
		fmt.Println(s, i)
	}
	fmt.Println("expected i = 4. real i =", i)
	// why is i=0? i expect it to be 4.
	c <- i
}

func main() {
	c := make(chan int)
	go say("Hello World", c)
	sts := <- c
	fmt.Println("c = ", sts)
}