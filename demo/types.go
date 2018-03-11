package demo

import (
	"fmt"
	"math/cmplx"
)

var (
	toBe   bool       = false
	Maxint uint       = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(4 + 5i)
)

func PrintType() {
	fmt.Printf("Type: %T value: %v\n", toBe, toBe)
	fmt.Printf("Type: %T value: %v\n", Maxint, Maxint)
	fmt.Printf("Type: %T value: %v\n", z, z)
}
