package main

import "fmt"

func main() {
	i := 1

	for i < 10 {
		fmt.Printf("loop %d\n", i)
		i++
	}

	fmt.Println("----------------------------------------")

	for i := 1; i < 10; i++ {
		fmt.Printf("loop %d\n", i)
	}
}
