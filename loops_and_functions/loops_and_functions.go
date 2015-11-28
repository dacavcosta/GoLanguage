package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (soma float64) {
	z := float64(2.)
	for {
		z = z - (math.Pow(z, 2) - x) / (2*z)
		if math.Abs(soma-z) < 1e-15 {
			break
		}
		soma = z
	}
	return
}

func main() {
	fmt.Println(Sqrt(4))
	fmt.Println(math.Sqrt(4))
}