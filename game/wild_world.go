package game

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/region"
	"golang.org/x/text/language"
)

const (
	// wildWorld is the common reference for Animal Crossing: Wild World.
	wildWorld = "WildWorld"
)

var (
	// wildWorldNameAmericanEnglish represents Animal Crossing: Wild World's name in American English.
	wildWorldNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Animal Crossing: Wild World",
	}
)

var (
	// wildWorldReleaseDateAustralia represents Animal Crossing: Wild World's release date in Australia.
	wildWorldReleaseDateAustralia = nook.ReleaseDate{
		Day:    8,
		Month:  time.December,
		Region: region.Australia,
		Year:   2005,
	}

	// wildWorldReleaseDateEurope represents Animal Crossing: Wild World's release date in Europe.
	wildWorldReleaseDateEurope = nook.ReleaseDate{
		Day:    31,
		Month:  time.March,
		Region: region.Europe,
		Year:   2006,
	}

	// wildWorldReleaseDateJapan represents Animal Crossing: Wild World's release date in Japan.
	wildWorldReleaseDateJapan = nook.ReleaseDate{
		Day:    23,
		Month:  time.November,
		Region: region.Japan,
		Year:   2005,
	}

	// wildWorldReleaseDateNorthAmerica represents Animal Crossing: Wild World's release date in North America.
	wildWorldReleaseDateNorthAmerica = nook.ReleaseDate{
		Day:    5,
		Month:  time.December,
		Region: region.NorthAmerica,
		Year:   2005,
	}
)

var (
	// wildWorldName contains the localized names of Animal Crossing: Wild World.
	wildWorldName = nook.Languages{
		language.AmericanEnglish: wildWorldNameAmericanEnglish,
	}
)

var (
	// wildWorldReleaseDates contains Animal Crossing: Wild World's regional release history in chronological order.
	wildWorldReleaseDates = []nook.ReleaseDate{
		wildWorldReleaseDateJapan,
		wildWorldReleaseDateNorthAmerica,
		wildWorldReleaseDateAustralia,
		wildWorldReleaseDateEurope,
	}
)

var (
	// WildWorld represents Animal Crossing: Wild World.
	WildWorld = nook.Game{
		Key:          nook.Key(wildWorld),
		Name:         wildWorldName,
		ReleaseDates: wildWorldReleaseDates,
		ReleaseOrder: 6,
	}
)
