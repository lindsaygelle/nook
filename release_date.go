package nook

import "time"

func compareReleaseDates(a, b ReleaseDate) int {
	switch {
	case a.Year < b.Year:
		return -1
	case a.Year > b.Year:
		return 1
	case a.Month < b.Month:
		return -1
	case a.Month > b.Month:
		return 1
	case a.Day < b.Day:
		return -1
	case a.Day > b.Day:
		return 1
	case a.Region.Key < b.Region.Key:
		return -1
	case a.Region.Key > b.Region.Key:
		return 1
	default:
		return 0
	}
}

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

// After reports whether the release date occurs after the provided release
// date.
func (r ReleaseDate) After(other ReleaseDate) bool {
	return r.Compare(other) > 0
}

// Before reports whether the release date occurs before the provided release
// date.
func (r ReleaseDate) Before(other ReleaseDate) bool {
	return r.Compare(other) < 0
}

// Compare compares two release dates by year, month, day, and then region key
// for deterministic ordering.
func (r ReleaseDate) Compare(other ReleaseDate) int {
	return compareReleaseDates(r, other)
}

// Ok returns a boolean indicating whether the ReleaseDate information is
// complete and valid.
func (r ReleaseDate) Ok() bool {
	return r.Day != 0 && r.Month != 0 && r.Region.Key != "" && r.Year != 0
}
