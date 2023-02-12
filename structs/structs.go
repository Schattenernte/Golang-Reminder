package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {

	p := person{name: name, age: 42}
	return &p
}

type employee struct {
	name string
	age  int
}

type Widget struct {
	id    int
	attrs []string
}

func main() {

	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Elena"}) // age otomatik 0 oldu

	fmt.Println(newPerson("John"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	// slice ile struct kullanma

	fmt.Printf("--------\n")

	employees := make([]employee, 2)

	employees[0] = employee{name: "ibrahim", age: 20}
	employees[1] = employee{name: "mehmet", age: 20}

	for _, e := range employees {
		fmt.Println(e)
	}

	// 2. kullanım

	widgets := []Widget{
		Widget{
			id:    10,
			attrs: []string{"blah", "foo"},
		},
		Widget{
			id:    20,
			attrs: []string{"foo", "bar"},
		},
		Widget{
			id:    20,
			attrs: []string{"xyz", "foo"},
		},
	}

	for _, j := range widgets {
		fmt.Print(j.id)
		for _, y := range j.attrs {
			fmt.Print(y)
		}
		fmt.Println()
	}

}
