package main

import (
	"fmt"
)

func slowGreet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func main() {
	done := make(chan bool)
	go slowGreet("How ... are ... you ...?", done)
	<- done
}