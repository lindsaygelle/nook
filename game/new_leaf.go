package game

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/region"
	"golang.org/x/text/language"
)

const (
	// newLeaf is the common reference for Animal Crossing: New Leaf.
	newLeaf = "NewLeaf"
)

var (
	// newLeafNameAmericanEnglish represents Animal Crossing: New Leaf's name in American English.
	newLeafNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Animal Crossing: New Leaf",
	}
)

var (
	// newLeafReleaseDateAustralia represents Animal Crossing: New Leaf's release date in Australia.
	newLeafReleaseDateAustralia = nook.ReleaseDate{
		Day:    15,
		Month:  time.June,
		Region: region.Australia,
		Year:   2013,
	}

	// newLeafReleaseDateEurope represents Animal Crossing: New Leaf's release date in Europe.
	newLeafReleaseDateEurope = nook.ReleaseDate{
		Day:    14,
		Month:  time.June,
		Region: region.Europe,
		Year:   2013,
	}

	// newLeafReleaseDateJapan represents Animal Crossing: New Leaf's release date in Japan.
	newLeafReleaseDateJapan = nook.ReleaseDate{
		Day:    8,
		Month:  time.November,
		Region: region.Japan,
		Year:   2012,
	}

	// newLeafReleaseDateNorthAmerica represents Animal Crossing: New Leaf's release date in North America.
	newLeafReleaseDateNorthAmerica = nook.ReleaseDate{
		Day:    9,
		Month:  time.June,
		Region: region.NorthAmerica,
		Year:   2013,
	}
)

var (
	// newLeafName contains the localized names of Animal Crossing: New Leaf.
	newLeafName = nook.Languages{
		language.AmericanEnglish: newLeafNameAmericanEnglish,
	}
)

var (
	// newLeafReleaseDates contains Animal Crossing: New Leaf's regional release history in chronological order.
	newLeafReleaseDates = []nook.ReleaseDate{
		newLeafReleaseDateJapan,
		newLeafReleaseDateNorthAmerica,
		newLeafReleaseDateEurope,
		newLeafReleaseDateAustralia,
	}
)

var (
	// NewLeaf represents Animal Crossing: New Leaf.
	NewLeaf = nook.Game{
		Key:          nook.Key(newLeaf),
		Name:         newLeafName,
		ReleaseDates: newLeafReleaseDates,
		ReleaseOrder: 8,
	}
)
