package clock

import (
	"fmt"
)

type Clock struct {
	minutes int
}

const dayMinutes = 24 * 60
func New(h, m int) Clock {
	minutes := (h*60 + m) % dayMinutes
	for minutes < 0 {
		minutes += dayMinutes
	}
	return Clock{minutes}
}

func (c Clock) Add(a int) Clock {
	return New(0, c.minutes+a)
}

func (c Clock) Subtract(a int) Clock {
	return New(0, c.minutes-a)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minutes/60, c.minutes%60)
}
