package main


import (
    "fmt"
    "time"
)

func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}

func main() {

    f("direct")

    go f("goroutine")

    go func(msg string) {
        fmt.Println(msg)
    }("going")

    time.Sleep(time.Second)
    fmt.Println("done")
}
/*
Goroutine, Go programlama dilinde çok hafif ve verimli iş parçacıklarıdır. 
Görevlerin (task) eşzamanlı bir şekilde çalıştırılması için kullanılırlar. Go dilinde iş parçacıkları, 
Goroutine'ler tarafından yönetilir ve bu iş parçacıkları, işlemci çekirdeklerine göre otomatik olarak çalıştırılır. 
Görevler arasındaki geçişler, işletim sistemi seviyesindeki iş parçacıkları arasındaki geçişlerden daha hızlı ve daha verimlidir.

Goroutine'ler, bir işlevin paralel olarak çalıştırılmasına olanak tanır. 
Görev, ana işlevin iş parçacıklarından biri tarafından çalıştırılmak üzere Goroutine tarafından oluşturulur ve bu işlev,
işlemci çekirdeği boşalana kadar veya belirtilen bir süre sona erene kadar çalışır. 
Görev tamamlandıktan sonra, Goroutine sonlanır ve kontrol tekrar ana iş parçacığına döner.

Goroutine'lerin özellikleri arasında çok hafif olmaları, bellek tüketimlerinin az olması ve çok sayıda Goroutine'in kullanılabilmesi gibi avantajlar bulunur. 
Bu nedenle, Go dilinde iş parçacığı bazlı programlama yerine Goroutine'ler kullanılarak eşzamanlılık sağlanması tercih edilir. 
Bu da Go dilinin hızlı ve verimli çalışmasına yardımcı olur.


*/