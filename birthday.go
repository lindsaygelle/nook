package nook

import (
	"time"
)

type Birthday struct {
	Day   uint8
	Month time.Month
}

func (v Birthday) Ok() bool {
	return (v.Day != 0) && (v.Month != 0)
}
