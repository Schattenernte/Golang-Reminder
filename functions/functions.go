package main

import "fmt"

func plusplus(a, b, c, d int) int {
	return a + b + c + d
}

func plus(a int, b int) int {
	return a + b
}

// multiple return functions

func vals() (int, int) {
	return 4, 8
}

// Variadic functions

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	// nums dizi olarak gittigi icin ilk _ o anki indeks boş bıraktık.
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	res := plus(2, 4)
	fmt.Println("2 + 4 = ", res)

	res = plusplus(2, 4, 8, 20)
	fmt.Println("2 + 4 + 8 + 20 = ", res)

	// multiple return functions

	a, b := vals()
	fmt.Println(a, b)

	// Variadic functions

	// sınırsız argüman alan fonksiyonlardaki sınırsız argümanlar
	// dizi şeklinde gider örneğin []int
	sum(2, 4)
	sum(2, 4, 6, 8)

	// bir dilimi de bu şekilde gönderebiliriz.
	nums := []int{2, 4, 6, 8}
	sum(nums...)
}
