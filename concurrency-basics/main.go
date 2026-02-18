package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello,", phrase)
	doneChan <- true
}

func longGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello, ", phrase)
	doneChan <- true
	close(doneChan)
}

func main() {
	done := make(chan bool)

	go greet("What are you doing?", done)

	go greet("How are you doing?", done)

	go longGreet("Nice to meet you how are you?", done)

	go greet("I hope you are doing good!", done)

	for range done {
	}

}
