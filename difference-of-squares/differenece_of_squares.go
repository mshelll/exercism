package diffsquares

import (
	//"fmt"
)

func SquareOfSum(in int) (out int) {

   sum := 0
   for i := 1 ; i <= in ; i++ {
      
	   sum += i
   }

   return sum*sum
}

func SumOfSquares(in int) (out int) {

	sum := 0
	for i := 1 ; i <= in ; i++ {
		
		sum += i*i
	}

	return sum
}

func Difference(in int) (out int) {

	return SquareOfSum(in) - SumOfSquares(in)
}
