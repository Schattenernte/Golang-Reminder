package main

import (
	"fmt"
	"time"
)

// done ifaesi done kanalından bir değer alana kadar programın beklenmesini sağladığı için
// done <- true yaptığımız gibi go routine harici main process programı sonlandırıyor bu yüzden 22. satırı sildiğimzde program hiçbir çıktı vermiyor
func worker(done chan bool) {
	fmt.Print("working..")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)

	<-done
}
