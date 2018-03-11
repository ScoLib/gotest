package demo

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	z := 10.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func TestLoops(v float64) {
	fmt.Printf("sqrt: %v\n", sqrt(v))
	fmt.Printf("math.Sqrt: %v\n", math.Sqrt(v))
}
