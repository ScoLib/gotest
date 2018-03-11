package demo

import (
	"math"
	"fmt"
)

func Convert() {
	x, y := 1, 2
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, f, z)
}
