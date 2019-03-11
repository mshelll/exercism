package clock

import (
	"fmt"
)

type Clock struct {
	Hour, Minute int
}


const MINUTE_OF_DAY = 24*60

func New(hours, minutes int) (out Clock) {

	var c1 Clock
	total_minutes := (MINUTE_OF_DAY + ((hours*60 + minutes)%MINUTE_OF_DAY))%MINUTE_OF_DAY
	c1.Hour = total_minutes/60
	c1.Minute = total_minutes%60
	return c1
}

func (c1 Clock) String() (out string) {

	return fmt.Sprintf("%02d:%02d", c1.Hour, c1.Minute)
}

func (c1 Clock) Add(minutes int) (out Clock) {

    total_minutes_clock := c1.Hour*60 + c1.Minute
	total_minutes := (total_minutes_clock + minutes)%MINUTE_OF_DAY
	c1.Hour = total_minutes/60
	c1.Minute = total_minutes%60
    return c1
}

func (c1 Clock) Subtract(minutes int) (out Clock) {

	total_minutes_clock := c1.Hour*60 + c1.Minute
	total_minutes := (MINUTE_OF_DAY + total_minutes_clock - minutes%MINUTE_OF_DAY)%MINUTE_OF_DAY
	c1.Hour = total_minutes/60
	c1.Minute = total_minutes%60
	return c1

}
