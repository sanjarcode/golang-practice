package main

import "fmt"

func main() {
str:="สวัสดีhello"
for i, chr:= range str {
	fmt.Print(i, string(chr), " ")
}
}
