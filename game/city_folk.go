package game

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/region"
	"golang.org/x/text/language"
)

const (
	// cityFolk is the common reference for Animal Crossing: City Folk.
	cityFolk = "CityFolk"
)

var (
	// cityFolkNameAmericanEnglish represents Animal Crossing: City Folk's name in American English.
	cityFolkNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Animal Crossing: City Folk",
	}
)

var (
	// cityFolkReleaseDateAustralia represents Animal Crossing: City Folk's release date in Australia.
	cityFolkReleaseDateAustralia = nook.ReleaseDate{
		Day:    4,
		Month:  time.December,
		Region: region.Australia,
		Year:   2008,
	}

	// cityFolkReleaseDateEurope represents Animal Crossing: City Folk's release date in Europe.
	cityFolkReleaseDateEurope = nook.ReleaseDate{
		Day:    5,
		Month:  time.December,
		Region: region.Europe,
		Year:   2008,
	}

	// cityFolkReleaseDateJapan represents Animal Crossing: City Folk's release date in Japan.
	cityFolkReleaseDateJapan = nook.ReleaseDate{
		Day:    20,
		Month:  time.November,
		Region: region.Japan,
		Year:   2008,
	}

	// cityFolkReleaseDateNorthAmerica represents Animal Crossing: City Folk's release date in North America.
	cityFolkReleaseDateNorthAmerica = nook.ReleaseDate{
		Day:    16,
		Month:  time.November,
		Region: region.NorthAmerica,
		Year:   2008,
	}
)

var (
	// cityFolkName contains the localized names of Animal Crossing: City Folk.
	cityFolkName = nook.Languages{
		language.AmericanEnglish: cityFolkNameAmericanEnglish,
	}
)

var (
	// cityFolkReleaseDates contains Animal Crossing: City Folk's regional release history in chronological order.
	cityFolkReleaseDates = []nook.ReleaseDate{
		cityFolkReleaseDateNorthAmerica,
		cityFolkReleaseDateJapan,
		cityFolkReleaseDateAustralia,
		cityFolkReleaseDateEurope,
	}
)

var (
	// CityFolk represents Animal Crossing: City Folk.
	CityFolk = nook.Game{
		Key:          nook.Key(cityFolk),
		Name:         cityFolkName,
		ReleaseDates: cityFolkReleaseDates,
		ReleaseOrder: 7,
	}
)
