package main

import (
	"errors"
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint(float64(e))
}

func Sqrt(x float64) (float64, error)  {
	if x < 0 {
		err := ErrNegativeSqrt(x)
		return 0, errors.New(err.Error())
	}
	return math.Sqrt(x), nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}