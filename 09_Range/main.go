package main

import (
	"fmt"
)

func main() {
	colors := []string{"red", "green", "blue", "purple", "orange"}

	for index, value := range colors {
		fmt.Printf("%d-%s\n", index, value)
	}

	fmt.Println("----------------------------------------")

	for _, value := range colors {
		fmt.Printf("value-%s\n", value)
	}

	fmt.Println("----------------------------------------")

	for index := range colors {
		fmt.Printf("index-%d\n", index)
	}

	fmt.Println("----------------------------------------")

	var romanNumberDict = map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}

	for key, value := range romanNumberDict {
		fmt.Printf("%d-%s\n", key, value)
	}

}
