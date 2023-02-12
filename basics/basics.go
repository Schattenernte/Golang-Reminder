package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("go" + "lang")
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	fmt.Println(!true)

	// basit değişkenler
	var a = "initial"
	fmt.Println(a)
	var b, c int = 1, 2
	fmt.Println(b, c)
	var e int
	fmt.Println(e) // otomatik olarak 0 başlatılır
	f := "apple"
	fmt.Print(f)

	fmt.Printf("\n---------------------------\n")
	// uint8 	0 to 255
	// uint16 	0 to 65535
	// uint32	0 to 4294967295
	// uint64	0 to 18446744073709551615

	// int8		-128 to 127
	// int16	-32768 to 32767
	// int32	-2147483648 to 2147483647
	// int64	-9223372036854775808 to 9223372036854775807

	// byte = int8 anlamında
	// uintptr = null pointer tanımlamak için

	// go döngüler

	// while
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// for

	for j := 6; j <= 8; j++ {
		fmt.Println(j)
	}

	// if else

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// switch case

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T", t)
		}
	}
	whatAmI(true)
	whatAmI("hey")

}
