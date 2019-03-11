package twelve

import (
	"fmt"
)

var days = [...]string{ "first", "second", "third", "fourth", "fifth", "sixth",
						"seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}

var elements = [...]string{
	"a Partridge in a Pear Tree.",
	"two Turtle Doves, and ",
	"three French Hens, ",
	"four Calling Birds, ",
	"five Gold Rings, ",
	"six Geese-a-Laying, ",
	"seven Swans-a-Swimming, ",
	"eight Maids-a-Milking, ",
	"nine Ladies Dancing, ",
	"ten Lords-a-Leaping, ",
	"eleven Pipers Piping, ",
	"twelve Drummers Drumming, ",
}

func fetch_prefix(day_count int) string {

	return fmt.Sprintf("On the %s day of Christmas my true love gave to me: ", days[day_count])
}

func fetch_element(day_count int) string {

	element := ""
	for i := day_count; i >= 0; i-- {
		element += elements[i]
	}
	return element
}

func fetch_line(prefix string, line_no int) string {

	return prefix + fetch_element(line_no)
}


func Song() string {

	song := ""
	for i:= 0; i < 12; i++ {

		prefix := fetch_prefix(i)
		song += prefix + fetch_element(i) + "\n"
	}
	return song
}

func Verse(line_no int) string {

	prefix := fetch_prefix(line_no-1)
	return prefix + fetch_element(line_no-1)
}