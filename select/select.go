package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
	}()
	// program select içerisinde herhangi bir case'e girdiğinde o case sonlanana kadar program durur c1 kanalından veri geldiği zaman c1 c2 den geldiğinide c2 çalışacak
	// eğer bir case bloğu işlem yapmak için hazır değilse ve diğer case bloklarına geçiş yapılırken,
	//önceki case bloğu daha sonra işlem yapmaya hazır hale gelirse, select ifadesi o case bloğuna geri dönecektir.
	// ^^
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
