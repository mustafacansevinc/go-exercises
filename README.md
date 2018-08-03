# go-exercises
Some practice and learning with go

## What is the purpose of this repo?
It just contains my learning process of using golang. Nothing more. Dont forget to run program as:
```
go run hello.go
```

Learn the difference between ```i := 0``` and ```i = 0```. first one gives and initial value, but what else? Is it for specifying type?

### goroutine
the keyword go means run as a seperate task. these are named goroutines.
when you run a function as goroutine, while it is being setup, the rest of the main() function continues.
by then the goroutine is active, and works.

### channels
goroutines can communicate using channels. now we can create a channel by sending a parameter to the method as
```
func say(s string, c chan string)
```

and in that method return the information with that channel as 
```
c <- s
```

finally, in the main function we have to create that channel with make command:
```
func main() {
c := make(chan string)
go say("Hello", c)
...
```
Here is an example program:
```
package main

import (
	"fmt"
)

func say(s string, c chan int) {
	var i int
	for i = 0; i < 5; i++ {
		fmt.Println(s, i)
	}
	c <- i
}

func main() {
	c := make(chan int)
	go say("Hello World", c)
	sts := <- c
	fmt.Println("c = ", sts)
}
```

the *fmt* package is a shorthand of "format".

