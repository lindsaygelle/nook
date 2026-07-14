package game

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/region"
	"golang.org/x/text/language"
)

const (
	// animalCrossing is the common reference for Animal Crossing.
	animalCrossing = "AnimalCrossing"
)

var (
	// animalCrossingNameAmericanEnglish represents Animal Crossing's name in American English.
	animalCrossingNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Animal Crossing",
	}
)

var (
	// animalCrossingReleaseDateAustralia represents Animal Crossing's release date in Australia.
	animalCrossingReleaseDateAustralia = nook.ReleaseDate{
		Day:    17,
		Month:  time.October,
		Region: region.Australia,
		Year:   2003,
	}

	// animalCrossingReleaseDateEurope represents Animal Crossing's release date in Europe.
	animalCrossingReleaseDateEurope = nook.ReleaseDate{
		Day:    24,
		Month:  time.September,
		Region: region.Europe,
		Year:   2004,
	}

	// animalCrossingReleaseDateNorthAmerica represents Animal Crossing's release date in North America.
	animalCrossingReleaseDateNorthAmerica = nook.ReleaseDate{
		Day:    16,
		Month:  time.September,
		Region: region.NorthAmerica,
		Year:   2002,
	}
)

var (
	// animalCrossingName contains the localized names of Animal Crossing.
	animalCrossingName = nook.Languages{
		language.AmericanEnglish: animalCrossingNameAmericanEnglish,
	}
)

var (
	// animalCrossingReleaseDates contains Animal Crossing's regional release history in chronological order.
	animalCrossingReleaseDates = []nook.ReleaseDate{
		animalCrossingReleaseDateNorthAmerica,
		animalCrossingReleaseDateAustralia,
		animalCrossingReleaseDateEurope,
	}
)

var (
	// AnimalCrossing represents Animal Crossing.
	AnimalCrossing = nook.Game{
		Key:          nook.Key(animalCrossing),
		Name:         animalCrossingName,
		ReleaseDates: animalCrossingReleaseDates,
		ReleaseOrder: 3,
	}
)
