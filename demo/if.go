package demo

import (
	"math"
	"fmt"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g > %g\n", v, lim)
	}

	return lim
}

func sqrtif(x float64) string {
	if x < 0 {
		return sqrtif(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func TestIf() {
	fmt.Println(sqrtif(2), sqrtif(-4))
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
}
