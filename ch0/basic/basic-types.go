package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe bool = false
	MaxInt uint64 = 1 << 64 -1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Values: %v\n\n", ToBe, ToBe)
	fmt.Printf("Type: %T Values: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Values: %v\n", z, z)
}