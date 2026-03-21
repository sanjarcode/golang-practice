package main

// imports must come before declarations
// multiple imports are kept in parentheses, no comma
// imports can be separated or in comma list
import (
	"fmt"
	"math"
)

const s string = "constant"

// non-declaration statement outside function body
// fmt.Println(s)

func main() {
	fmt.Println(1, 2, 3)

	const n float32 = 500000000
	const k = 500000001

	fmt.Println(k == n)

	const d float32 = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(math.Pi))
}
