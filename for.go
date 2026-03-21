package main

import "fmt"

func main() {

	// range x is a thing

	// braces are needed in break and continue

	// for i:= range 10 {
	// 	fmt.Println(i)
	// }

	// for i:=0; i < 20; i++ {
	// 	fmt.Println(i)
	// }

	for i := 0; i < 4; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
