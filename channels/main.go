package main

import (
	"fmt"
	"time"
)

func main() {
	msg := make(chan string) // channel declaration with type string
	go greet(msg)            // start new goroutine

	time.Sleep(2 * time.Second)

	greeting := <-msg // receives the message from the channel

	time.Sleep(2 * time.Second)
	fmt.Println("Greeting received")
	fmt.Println(greeting)

	_, ok := <-msg // check if channels is open/closed
	if ok {
		fmt.Println("Channel is open!")
	} else {
		fmt.Println("Channel is closed!")
	}
}

func greet(ch chan string) {
	fmt.Println("Greeter waiting to send greeting!")

	ch <- "Hello World"
	close(ch) // close the channel

	fmt.Println("Greeter completed")
}
