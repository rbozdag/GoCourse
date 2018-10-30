package main

import "fmt"

func main() {
	printComparitionResult(2, 10)
	printColor("0x00ff00")
}

func printComparitionResult(x, y int) {
	if x == y {
		fmt.Printf("%d equals %d\n", x, y)
	} else if x > y {
		fmt.Printf("%d is less than %d\n", y, x)
	} else {
		fmt.Printf("%d is less than %d\n", x, y)
	}
}

func printColor(colorCode string) {
	switch colorCode {
	case "0xff0000":
		println("color is red")
	case "0x00ff00":
		println("color is green")
	case "0x0000ff":
		println("color is blue")
	default:
		println("color is mixed")
	}
}
