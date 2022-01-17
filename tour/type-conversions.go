package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 13, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)

	fmt.Println(x, y, f, z)

	negative := -42
	un := uint(negative)
	fmt.Println(negative, un)
}
