package main

import "fmt"

func main() {
	mp:=map[string]int {"hello": 1, "world": 2}
	for k, v := range mp {
		fmt.Println(k, v)
	}

	for c:= range "go" {
		fmt.Println(string(c))
	}
}
