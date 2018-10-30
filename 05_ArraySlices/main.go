package main

import (
	"fmt"
)

func main() {
	var fruitArr [2]string
	fruitArr[0] = "Apple"
	fruitArr[1] = "Lemon"
	fmt.Println("fruitArr>>", fruitArr)
	fmt.Println("fruitArr[0]>>", fruitArr[0])

	platforms := [2]string{"IOS", "Android"}
	fmt.Println("platforms>>", platforms)

	fruitSlice := []string{"Grape", "Banana"}
	fruitSlice = append(fruitSlice, "Lemon")
	fmt.Println("fruitSlice>>", fruitSlice)
	fmt.Println("fruitSlice.len>>", len(fruitSlice))
}
