package main

import "fmt"

type rect struct {
	width, height int
}

// pointer ile de tanımlanabilir
func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area : ", r.area())
	fmt.Println("perim : ", r.perim())

	rp := &r // & ile atama yaparsak r yi veya rp yi değiştirdiğmizde ikisi de değişir normal de atanabili
	fmt.Println("area : ", rp.area())
	fmt.Println("perim:", rp.perim())
}
