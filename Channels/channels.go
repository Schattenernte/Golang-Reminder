package main

import "fmt"

// kanallar goroutineler arası veri transferini sağlar

func main() {
	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)

}
