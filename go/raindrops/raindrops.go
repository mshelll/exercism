package raindrops


import (
	"strings"
	"strconv"
)

func Convert(n int) (s string) {

	var op []string

	if n % 3 == 0 {

		op = append(op, "Pling")
	}
	if n % 5 == 0 {

		op = append(op, "Plang")
	}
	if n % 7 == 0 {

		op = append(op, "Plong")
	}

	if len(op) == 0 {

		op = append(op, strconv.Itoa(n))
	}

	return strings.Join(op, "")
}