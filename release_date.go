package nook

import "time"

// ReleaseDate represents a game's release date in a specific region.
type ReleaseDate struct {
	// Day is the day of the month when the game released.
	Day uint8

	// Month is the month of the year when the game released.
	Month time.Month

	// Region is the major region where the game released on this date.
	Region Region

	// Year is the year when the game released.
	Year uint16
}

// Ok returns a boolean indicating whether the ReleaseDate information is
// complete and valid.
func (r ReleaseDate) Ok() bool {
	return r.Day != 0 && r.Month != 0 && r.Region.Key != "" && r.Year != 0
}
