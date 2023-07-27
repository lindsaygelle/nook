package nook

import (
	"time"
)

// Birthday represents the birthday information of an Animal Crossing character.
type Birthday struct {
	// Day is the day of the month when the character's birthday falls.
	Day uint8

	// Month is the month of the year when the character's birthday falls.
	Month time.Month
}

// Ok returns a boolean indicating whether the Birthday information is complete and valid.
// It checks if both the Day and Month fields have been set.
func (b Birthday) Ok() bool {
	return (b.Day != 0) && (b.Month != 0)
}
