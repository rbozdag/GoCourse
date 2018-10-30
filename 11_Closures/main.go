package main

import (
	"fmt"
)

func main() {
	var sum = adder()
	fmt.Printf("%T\n", sum)
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", sum(i))
	}

	fmt.Println("----------------------------------------")

	sum = adder2()
	fmt.Printf("%T\n", sum)
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", sum(i))
	}
}

func adder() func(int) int {
	var sum = 0
	return func(val int) int {
		sum = sum + val
		return sum
	}
}

func adder2() func(int) int {
	var sum = 0
	var retval func(val int) int
	retval = func(val int) int {
		sum = sum + val
		return sum
	}

	return retval
}
