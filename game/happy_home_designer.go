package game

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/region"
	"golang.org/x/text/language"
)

const (
	// happyHomeDesigner is the common reference for Animal Crossing: Happy Home Designer.
	happyHomeDesigner = "HappyHomeDesigner"
)

var (
	// happyHomeDesignerNameAmericanEnglish represents Animal Crossing: Happy Home Designer's name in American English.
	happyHomeDesignerNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Animal Crossing: Happy Home Designer",
	}
)

var (
	// happyHomeDesignerReleaseDateAustralia represents Animal Crossing: Happy Home Designer's release date in Australia.
	happyHomeDesignerReleaseDateAustralia = nook.ReleaseDate{
		Day:    3,
		Month:  time.October,
		Region: region.Australia,
		Year:   2015,
	}

	// happyHomeDesignerReleaseDateEurope represents Animal Crossing: Happy Home Designer's release date in Europe.
	happyHomeDesignerReleaseDateEurope = nook.ReleaseDate{
		Day:    2,
		Month:  time.October,
		Region: region.Europe,
		Year:   2015,
	}

	// happyHomeDesignerReleaseDateJapan represents Animal Crossing: Happy Home Designer's release date in Japan.
	happyHomeDesignerReleaseDateJapan = nook.ReleaseDate{
		Day:    30,
		Month:  time.July,
		Region: region.Japan,
		Year:   2015,
	}

	// happyHomeDesignerReleaseDateNorthAmerica represents Animal Crossing: Happy Home Designer's release date in North America.
	happyHomeDesignerReleaseDateNorthAmerica = nook.ReleaseDate{
		Day:    25,
		Month:  time.September,
		Region: region.NorthAmerica,
		Year:   2015,
	}
)

var (
	// happyHomeDesignerName contains the localized names of Animal Crossing: Happy Home Designer.
	happyHomeDesignerName = nook.Languages{
		language.AmericanEnglish: happyHomeDesignerNameAmericanEnglish,
	}
)

var (
	// happyHomeDesignerReleaseDates contains Animal Crossing: Happy Home Designer's regional release history in chronological order.
	happyHomeDesignerReleaseDates = []nook.ReleaseDate{
		happyHomeDesignerReleaseDateJapan,
		happyHomeDesignerReleaseDateNorthAmerica,
		happyHomeDesignerReleaseDateEurope,
		happyHomeDesignerReleaseDateAustralia,
	}
)

var (
	// HappyHomeDesigner represents Animal Crossing: Happy Home Designer.
	HappyHomeDesigner = nook.Game{
		Key:          nook.Key(happyHomeDesigner),
		Name:         happyHomeDesignerName,
		ReleaseDates: happyHomeDesignerReleaseDates,
		ReleaseOrder: 9,
	}
)
