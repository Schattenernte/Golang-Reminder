package main

import "fmt"

func main() {
	// slices usage
	nums := []int{2, 4, 8}
	sum := 0

	// _ indeksleri tutuyor num ise indeksteki değerleri indekse ihtiyacımız olmadığı için boş tanımladık
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum :", sum)

	// maps usage
	kvs := map[string]string{"a" : "apple", "b" : "banana"}
	for k,v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key :", k)
	}

	for _,k := range kvs {
		fmt.Println("value :",k)
	}
}
