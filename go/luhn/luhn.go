package luhn

import (
   "fmt"
   "unicode"
)

func Valid(in string) (out bool) {

	//var valid bool

	fmt.Printf("\n Input :%s", in)

	if len(in) <= 1 {

		//valid := false
		return false 
	}

	//numDigits := len(in)

	isSecond := true
	sum := 0

	for i := len(in) - 1 ; i >= 0 ; i-- {

		char := rune(in[i])

		if char == ' ' || char == '-' {
			continue
		}else if unicode.IsDigit(char) {

			d := int(char - '0')
			if isSecond {
				if d*2 > 9 {
					d = d*2 - 9
				} else {
					d = d*2
				}
			}
			fmt.Printf("\n d : %d", d)
			sum += d
		}else {
			return false
		}
		isSecond = !isSecond
	}

	fmt.Printf("\n Sum : %d", sum)
	return (sum%10 == 0)

	// var credit_number []int
	// for _, char := range in {

	// 	if char == ' ' {
	// 		continue
	// 	}else {
	// 		num := int(char-'0')
	// 		credit_number = append(credit_number, num)
	// 	}
	// }

	// fmt.Println(credit_number)

	// for index, num := range credit_number {

    //     if index%2 == 0 {
    //        num = num * 2
	// 	   if num >= 9 {
	// 		   num = num - 9
	// 	   }
    //        credit_number[index] = num 
	// 	}
	// }

	// sum := 0
	// for _, val := range credit_number  {

	// 	sum += val	
	// }

	// return sum%10 == 0
}