// package main

// import "fmt"

// func zeroval(i int) {
//     // default is call by value
// 	i = 0;
// }

// func zeroptr(iptr *int) {
// 	// call by reference explicit
// 	*iptr = 0;
// }

// func complex(arr []int,mp map[string]int) {
// 	// arrays passed by references
// 	mp["new_message"] = -1
// 	arr[0] = -1
// }


// func main() {
// 	a:=2
// 	fmt.Println("initial", a)

// 	zeroval(a)
// 	fmt.Println("zeroval", a)

// 	zeroptr(&a)
// 	fmt.Println("zeroptr", a)

// 	arr:=[]int {1, 2, 3}
// 	mp:=map[string]int {"hello": 1, "world": 2}

// 	fmt.Println("initial", arr, mp)
// 	complex(arr, mp)
// 	fmt.Println("after complex", arr, mp)
// }


package main

import "fmt"

func ops(arr []int,mp map[string]int) {
	// arrays passed by references
	mp["new_message"] = -1
	arr[0] = -1
}


func main() {
	arr:=[]int {1, 2, 3}
	mp:=map[string]int {"hello": 1, "world": 2}
	msg:="hello"

	fmt.Println("initial", arr, mp, msg)
	ops(arr, mp)
	fmt.Println("after complex", arr, mp, msg[0])
}

