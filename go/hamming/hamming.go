package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {
	
	count := 0

	if len(a) == len(b) {

        for i, _ := range(a) {

             if a[i] != b[i] {

				 count += 1
			 }
		}

	} else {
		return -1, errors.New("unequal length")
	}

	return count, nil
}
