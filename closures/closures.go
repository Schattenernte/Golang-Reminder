package main

import "fmt"

// geriye dönen değeri bir değişkene eşitlediğimizde fonksiyonun ana kısmındaki değişkenler statik tanımlanmış gibi olur yani o fonksiyonu eşitlediğimiz değişken tekrar çağrıldığında
// fonksiyonun içindeki değişken değerleri aynı kalır
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// geriye dönen fonksiyonun alacağı parametreleri adder'in yanındaki func un içine yazdık bu genel bir go syntaxı
// return kısmında da ona uygun bir şekilde parametreleri verdik (y yi kullanmadım gösterme amaçlı yazdım)
func adder() func(int, int) int {
	sum := 0
	return func(x int, y int) int {
		sum += x
		return sum
	}
}


func main() {
	// intseq'in geriye döndürdüğü şey bir fonksiyon olduğu için ona eşitlediğimiz şey artık bir fonksiyon oluyor
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
	fmt.Println(newInts())

	fmt.Printf("--------------------\n")
	
	adder := adder()
	fmt.Println(adder(2, 4))
	fmt.Println(adder(2, 4))
	fmt.Println(adder(2, 4))
	fmt.Println(adder(2, 4))
}
