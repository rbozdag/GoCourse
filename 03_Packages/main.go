package main

import (
	"fmt"
	"math"

	"gitlab.com/bozdag.rahmi/GoCourse/03_Packages/strutil"
)

func main() {
	fNumber := 2.6
	fmt.Println(math.Floor(fNumber))
	fmt.Println(strutil.Reverse("Hello"))
}
