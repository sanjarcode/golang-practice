package main

import "fmt"

func getCounter() func() int {
	state := 0
	return func() (int) {
		state+=1
		return state
	}
}

func getCounterCustom() func(int) int {
	state := 0
	return func(addendum int) (x int) {
		state += addendum
		return state
	}
}

func main() {
	counter := getCounter()

	fmt.Println(counter())
	fmt.Println(counter())

	counter2 := getCounter()
	fmt.Println(counter2())

	counterCustom := getCounterCustom()
	fmt.Println(counterCustom(1))
	fmt.Println(counterCustom(10))

}
