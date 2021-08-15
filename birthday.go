package nook

import (
	"time"
)

// Birthday is the Animal Crossing characters birthday information.
type Birthday struct {
	Day   uint8
	Month time.Month
}

func (v Birthday) Ok() bool {
	return (v.Day != 0) && (v.Month != 0)
}
