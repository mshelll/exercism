package grains

import (
	"errors"
	"math"
)

func Square(n int) (out uint64, err error) {

	if n >=1 && n <= 64 {
		return uint64(math.Pow(2, float64(n-1))), nil
	}

	return 0, errors.New("invalid")
}


func Total() (out uint64) {

    sum := uint64(0)
	for i:=1; i<=64; i++ {
		square, _ := Square(i)
		sum += square
	}

	return sum
}