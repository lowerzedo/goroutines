package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2) // number of goroutines to wait
	go helloworld(&wg)
	go goodbye(&wg)
	wg.Wait() // blocks exec untill internal counter == 0
}

func helloworld(wg *sync.WaitGroup) {
	defer wg.Done() // reduces internal counter by 1
	fmt.Println("Hello World!")
}

func goodbye(wg *sync.WaitGroup) {
	defer wg.Done() // reduces internal counter by 1
	fmt.Println("Good Bye!")
}
