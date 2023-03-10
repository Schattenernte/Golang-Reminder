package main

import (
	"fmt"
	"time"
)

func worker1(ch chan string) {
	// Kanala veri gönderme
	ch <- "Hello from worker1!"
}

func worker2(ch chan string) {
	// Kanaldan veri alıp ekrana yazdırma
	msg := <-ch
	fmt.Println(msg)
}

func main() {
	// Kanal oluşturma
	ch := make(chan string)

	// İki farklı go rutini oluşturma ve kanalı paylaştırma
	go worker1(ch)
	go worker2(ch)

	// Programın sonlanmasını beklemek için bekleyici kullanma
	time.Sleep(time.Second)
}
