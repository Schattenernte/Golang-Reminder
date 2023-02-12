package main

import "fmt"

func main() {
	// arrays

	var a [6]int
	fmt.Println(a)
	a[4] = 100
	fmt.Println("set : ", a)
	fmt.Println("get : ", a[4])
	fmt.Println("len : ", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl :", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d : ", twoD)

	fmt.Printf("\n--------\n")
	// Slices //////////////////////////////////////////

	s := make([]string, 3) // boş bir slice tanımlamak için make kullandık
	fmt.Println("emp : ", s)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set :", s)
	fmt.Println("get :", s[2])
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd :", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy :", c)

	l := s[2:5]
	fmt.Println("sl :", l)

	l = s[2:]
	fmt.Println("sl2", l)

	t := []string{"g", "h", "i"} // tek satırda değişken atama
	fmt.Println("dcl :", t)

	twoDD := make([][]int, 3) // [[] [] []]  ilk [] hepsini tutan bir kapsayıcı ikinci [] onun içindeki 3 tanesi
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoDD[i] = make([]int, innerLen) // ikinci üç kapsayıcıyı kullanmak istiyorsak kullanmak istediklerimizin her birini make ile oluşturmamız lazım 2. parametre o dizinin içinde kaç eleman olacağı
		for j := 0; j < innerLen; j++ {
			twoDD[i][j] = i + j
		}
	}
	fmt.Println("2d :", twoDD)

	db := make([][]int, 3)
	db[0] = make([]int, 3)
	db[1] = make([]int, 2)
	db[0][1] = 1
	db[1][1] = 4
	fmt.Println("a :", db)

	// Maps /////////////////////////////////////////////

	m := make(map[string]int)
	m["k1"] = 4
	m["k2"] = 8
	fmt.Println("map :", m)
	delete(m, "k2")

	n := map[string]int{"foo" : 1, "bar" : 2}
	fmt.Println("map : ", n)
	
}
