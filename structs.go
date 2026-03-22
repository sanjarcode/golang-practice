package main

import "fmt"

// declaring a struct
type node struct {
	value int
	next  *node
}

func main() {
	head := node{value: 1, next: nil} // initialization
	head.next = &node{value: 2, next: nil}
	head.next.next = &node{value: 7, next: nil}

	trav := &head
	for trav != nil {
		fmt.Println((*trav).value)
		trav = (*trav).next
	}

	str:= struct {
		name   string
		isGood bool
	}
	fmt.Println(type(str){"Tommy", false})
}
