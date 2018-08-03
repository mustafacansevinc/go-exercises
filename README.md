# go-exercises
Some practice and learning with go

## What is the purpose of this repo?
It just contains my learning process of using golang. Nothing more.

### goroutine
the keyword go means run as a seperate task. these are named goroutines.
when you run a function as goroutine, while it is being setup, the rest of the main() function continues.
by then the goroutine is active, and works.

### channels
goroutines can communicate using channels. now we can create a channel by sending a parameter to the method as
```
func say(s string, c chan int)
```

and in that method return the information with that channel as 
```
c <- s
```

...
