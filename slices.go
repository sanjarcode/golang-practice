package main

import (
	"fmt"
)

func main() {

	// x :=[5]int {1, 2, 3, 4}
	// x = append(x, 5)

	// fmt.Println(x, x[:3])
	// var s []string
	// fmt.Println("uninit:", s, s == nil, len(s) == 0)

	a := [...]string{"hello", "world", "bingo", "ones", "twos"}
	s := a[1:4]
	s = append(s, "three")
	fmt.Println("a", a, s)
	// s := make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// c := s[:3]
	// fmt.Println("c", c)

	s[0] = "a" // materializing a slice
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[1] = "b"
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[2] = "c"
	// c[2] = "d"
	fmt.Println("set:", s)
	// fmt.Println("c", c)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := s
	s[0] = "hello"
	fmt.Println("cpy:", c, s)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// t2 := []string{"g", "h", "i"}
	// if slices.Equal(t, t2) {
	// 	fmt.Println("t == t2")
	// }

	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

func AppendByte(slice []int, data ...int) []int {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]int, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}
