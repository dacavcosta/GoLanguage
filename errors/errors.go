package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string{
	return fmt.Sprintf("cannot Sqrt negative number: %g",
		float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0{
		return 0, ErrNegativeSqrt(x)
	}
	z := float64(2.)
	soma := float64(0)
	for {
		z = z - (math.Pow(z, 2) - x) / (2*z)
		if math.Abs(soma-z) < 1e-15 {
			break
		}
		soma = z
	}
	return soma, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}