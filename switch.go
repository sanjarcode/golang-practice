package main

import (
	"fmt"
	"time"
)

func whatAmI(i interface{}) {
	t := i.(type)
	switch t {
	case bool:
		fmt.Println("I'm a bool")
	case int:
		fmt.Println("I'm an int")
	default:
		fmt.Printf("Don't know type %T\n", t)
	}
}
func main() {

	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	j := time.Now().Weekday()
	fmt.Printf("Variable: name, Value: %v, Type: %t\n", j, j)
	switch j {
	case time.Sunday, time.Saturday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	z := t
	switch {
	case t == z:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
