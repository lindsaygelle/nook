package game

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/region"
	"golang.org/x/text/language"
)

const (
	// amiiboFestival is the common reference for Animal Crossing: amiibo Festival.
	amiiboFestival = "AmiiboFestival"
)

var (
	// amiiboFestivalNameAmericanEnglish represents Animal Crossing: amiibo Festival's name in American English.
	amiiboFestivalNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Animal Crossing: amiibo Festival",
	}
)

var (
	// amiiboFestivalReleaseDateAustralia represents Animal Crossing: amiibo Festival's release date in Australia.
	amiiboFestivalReleaseDateAustralia = nook.ReleaseDate{
		Day:    21,
		Month:  time.November,
		Region: region.Australia,
		Year:   2015,
	}

	// amiiboFestivalReleaseDateEurope represents Animal Crossing: amiibo Festival's release date in Europe.
	amiiboFestivalReleaseDateEurope = nook.ReleaseDate{
		Day:    21,
		Month:  time.November,
		Region: region.Europe,
		Year:   2015,
	}

	// amiiboFestivalReleaseDateJapan represents Animal Crossing: amiibo Festival's release date in Japan.
	amiiboFestivalReleaseDateJapan = nook.ReleaseDate{
		Day:    13,
		Month:  time.November,
		Region: region.Japan,
		Year:   2015,
	}

	// amiiboFestivalReleaseDateNorthAmerica represents Animal Crossing: amiibo Festival's release date in North America.
	amiiboFestivalReleaseDateNorthAmerica = nook.ReleaseDate{
		Day:    20,
		Month:  time.November,
		Region: region.NorthAmerica,
		Year:   2015,
	}
)

var (
	// amiiboFestivalName contains the localized names of Animal Crossing: amiibo Festival.
	amiiboFestivalName = nook.Languages{
		language.AmericanEnglish: amiiboFestivalNameAmericanEnglish,
	}
)

var (
	// amiiboFestivalReleaseDates contains Animal Crossing: amiibo Festival's regional release history in chronological order.
	amiiboFestivalReleaseDates = []nook.ReleaseDate{
		amiiboFestivalReleaseDateJapan,
		amiiboFestivalReleaseDateNorthAmerica,
		amiiboFestivalReleaseDateAustralia,
		amiiboFestivalReleaseDateEurope,
	}
)

var (
	// AmiiboFestival represents Animal Crossing: amiibo Festival.
	AmiiboFestival = nook.Game{
		Key:          nook.Key(amiiboFestival),
		Name:         amiiboFestivalName,
		ReleaseDates: amiiboFestivalReleaseDates,
		ReleaseOrder: 10,
	}
)
