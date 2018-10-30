package main

import (
	"fmt"
)

func main() {
	emailMap := make(map[string]string)
	emailMap["rahmi"] = "bozdag.rahmi@gmail.com"
	emailMap["bob"] = "bob3453453@gmail.com"

	fmt.Printf("full map: %s\n", emailMap)
	fmt.Printf("map[rahmi]: %s\n", emailMap["rahmi"])

	fmt.Println("----------------------------------------")

	delete(emailMap, "bob")
	fmt.Printf("full map after delete bob: %s\n", emailMap)

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

	fmt.Printf("Roman Number representaion of 1000 is %s\n", romanNumberDict[1000])
}
