package main

import "fmt"

// structs can have methods, these are . functions
// In go, these are added outside the struct but have a param of sorts
type rect struct {
	width, height int
}

func (r rect) area() int {
	return r.width * r.height
}

// you can accept pointers
func (rp *rect) perim() int {
	return 2 * (rp.width + rp.height)
}

func main() {
	r := rect{width: 10, height: 5}
	fmt.Println(r.area())
	fmt.Println(r.perim())

	rp:=&r;
	fmt.Println(rp.area())
	fmt.Println(rp.perim())

}
