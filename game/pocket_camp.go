package game

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/region"
	"golang.org/x/text/language"
)

const (
	// pocketCamp is the common reference for Animal Crossing: Pocket Camp.
	pocketCamp = "PocketCamp"
)

var (
	// pocketCampNameAmericanEnglish represents Animal Crossing: Pocket Camp's name in American English.
	pocketCampNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Animal Crossing: Pocket Camp",
	}
)

var (
	// pocketCampReleaseDateAustralia represents Animal Crossing: Pocket Camp's release date in Australia.
	pocketCampReleaseDateAustralia = nook.ReleaseDate{
		Day:    25,
		Month:  time.October,
		Region: region.Australia,
		Year:   2017,
	}

	// pocketCampReleaseDateWorldwide represents Animal Crossing: Pocket Camp's worldwide release date.
	pocketCampReleaseDateWorldwide = nook.ReleaseDate{
		Day:    21,
		Month:  time.November,
		Region: region.Worldwide,
		Year:   2017,
	}
)

var (
	// pocketCampName contains the localized names of Animal Crossing: Pocket Camp.
	pocketCampName = nook.Languages{
		language.AmericanEnglish: pocketCampNameAmericanEnglish,
	}
)

var (
	// pocketCampReleaseDates contains Animal Crossing: Pocket Camp's regional release history in chronological order.
	pocketCampReleaseDates = []nook.ReleaseDate{
		pocketCampReleaseDateAustralia,
		pocketCampReleaseDateWorldwide,
	}
)

var (
	// PocketCamp represents Animal Crossing: Pocket Camp.
	PocketCamp = nook.Game{
		Key:          nook.Key(pocketCamp),
		Name:         pocketCampName,
		ReleaseDates: pocketCampReleaseDates,
		ReleaseOrder: 11,
	}
)
