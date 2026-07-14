package game

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/platform"
	"github.com/lindsaygelle/nook/region"
	"golang.org/x/text/language"
)

const (
	// newHorizons is the common reference for Animal Crossing: New Horizons.
	newHorizons = "NewHorizons"
)

var (
	// newHorizonsNameAmericanEnglish represents Animal Crossing: New Horizons's name in American English.
	newHorizonsNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Animal Crossing: New Horizons",
	}
)

var (
	// newHorizonsReleaseDateWorldwide represents Animal Crossing: New Horizons' worldwide release date.
	newHorizonsReleaseDateWorldwide = nook.ReleaseDate{
		Day:    20,
		Month:  time.March,
		Region: region.Worldwide,
		Year:   2020,
	}
)

var (
	// newHorizonsName contains the localized names of Animal Crossing: New Horizons.
	newHorizonsName = nook.Languages{
		language.AmericanEnglish: newHorizonsNameAmericanEnglish,
	}
)

var (
	// newHorizonsPlatforms contains Animal Crossing: New Horizons' release platforms in deterministic key order.
	newHorizonsPlatforms = []nook.Platform{
		platform.NintendoSwitch,
	}
)

var (
	// newHorizonsReleaseDates contains Animal Crossing: New Horizons' regional release history in chronological order.
	newHorizonsReleaseDates = []nook.ReleaseDate{
		newHorizonsReleaseDateWorldwide,
	}
)

var (
	// NewHorizons represents Animal Crossing: New Horizons.
	NewHorizons = nook.Game{
		Key:          nook.Key(newHorizons),
		Name:         newHorizonsName,
		Platforms:    newHorizonsPlatforms,
		ReleaseDates: newHorizonsReleaseDates,
		ReleaseOrder: 12,
	}
)
