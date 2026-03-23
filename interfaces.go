package main

import ("fmt"
"math"
)

// represent the interface, so collected
// has to be told before
type geometry interface {
	area() float64
	perim() float64
}


type rect struct {
	width, height float64
}

// now supposes shape circle also needs area and perim
// naively, one would copy the methods
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

// you can accept pointers
func (rp rect) perim() float64 {
	return 2 * (rp.width + rp.height)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// you can accept pointers
func (cp circle) perim() float64 {
	return 2 * math.Pi * cp.radius
}


func measure(g geometry) {
	// g can be any interface realization
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func detectCircle(g geometry) {
	if c, ok:= g.(circle); ok {
		fmt.Println("circle with radius", c.radius);
	}
}

func main() {
	r := rect{width: 3, height: 4}
    c := circle{radius: 5}

    measure(r)
    measure(c)
    detectCircle(r)
    detectCircle(c)
}
