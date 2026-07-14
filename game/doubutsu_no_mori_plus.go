package game

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/platform"
	"github.com/lindsaygelle/nook/region"
	"golang.org/x/text/language"
)

const (
	// doubutsuNoMoriPlus is the common reference for Doubutsu no Mori+.
	doubutsuNoMoriPlus = "DoubutsuNoMoriPlus"
)

var (
	// doubutsuNoMoriPlusNameAmericanEnglish represents Doubutsu no Mori+'s name in American English.
	doubutsuNoMoriPlusNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Doubutsu no Mori+",
	}
)

var (
	// doubutsuNoMoriPlusReleaseDateJapan represents Doubutsu no Mori+'s release date in Japan.
	doubutsuNoMoriPlusReleaseDateJapan = nook.ReleaseDate{
		Day:    14,
		Month:  time.December,
		Region: region.Japan,
		Year:   2001,
	}
)

var (
	// doubutsuNoMoriPlusName contains the localized names of Doubutsu no Mori+.
	doubutsuNoMoriPlusName = nook.Languages{
		language.AmericanEnglish: doubutsuNoMoriPlusNameAmericanEnglish,
	}
)

var (
	// doubutsuNoMoriPlusPlatforms contains Doubutsu no Mori+'s release platforms in deterministic key order.
	doubutsuNoMoriPlusPlatforms = []nook.Platform{
		platform.GameCube,
	}
)

var (
	// doubutsuNoMoriPlusReleaseDates contains Doubutsu no Mori+'s regional release history in chronological order.
	doubutsuNoMoriPlusReleaseDates = []nook.ReleaseDate{
		doubutsuNoMoriPlusReleaseDateJapan,
	}
)

var (
	// DoubutsuNoMoriPlus represents Doubutsu no Mori+.
	DoubutsuNoMoriPlus = nook.Game{
		Key:          nook.Key(doubutsuNoMoriPlus),
		Name:         doubutsuNoMoriPlusName,
		Platforms:    doubutsuNoMoriPlusPlatforms,
		ReleaseDates: doubutsuNoMoriPlusReleaseDates,
		ReleaseOrder: 2,
	}
)
