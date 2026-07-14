package game

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/platform"
	"github.com/lindsaygelle/nook/region"
	"golang.org/x/text/language"
)

const (
	// doubutsuNoMoriEPlus is the common reference for Doubutsu no Mori e+.
	doubutsuNoMoriEPlus = "DoubutsuNoMoriEPlus"
)

var (
	// doubutsuNoMoriEPlusNameAmericanEnglish represents Doubutsu no Mori e+'s name in American English.
	doubutsuNoMoriEPlusNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Doubutsu no Mori e+",
	}
)

var (
	// doubutsuNoMoriEPlusReleaseDateJapan represents Doubutsu no Mori e+'s release date in Japan.
	doubutsuNoMoriEPlusReleaseDateJapan = nook.ReleaseDate{
		Day:    27,
		Month:  time.June,
		Region: region.Japan,
		Year:   2003,
	}
)

var (
	// doubutsuNoMoriEPlusName contains the localized names of Doubutsu no Mori e+.
	doubutsuNoMoriEPlusName = nook.Languages{
		language.AmericanEnglish: doubutsuNoMoriEPlusNameAmericanEnglish,
	}
)

var (
	// doubutsuNoMoriEPlusPlatforms contains Doubutsu no Mori e+'s release platforms in deterministic key order.
	doubutsuNoMoriEPlusPlatforms = []nook.Platform{
		platform.GameCube,
	}
)

var (
	// doubutsuNoMoriEPlusReleaseDates contains Doubutsu no Mori e+'s regional release history in chronological order.
	doubutsuNoMoriEPlusReleaseDates = []nook.ReleaseDate{
		doubutsuNoMoriEPlusReleaseDateJapan,
	}
)

var (
	// DoubutsuNoMoriEPlus represents Doubutsu no Mori e+.
	DoubutsuNoMoriEPlus = nook.Game{
		Key:          nook.Key(doubutsuNoMoriEPlus),
		Name:         doubutsuNoMoriEPlusName,
		Platforms:    doubutsuNoMoriEPlusPlatforms,
		ReleaseDates: doubutsuNoMoriEPlusReleaseDates,
		ReleaseOrder: 4,
	}
)
