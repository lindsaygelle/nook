package game

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/platform"
	"github.com/lindsaygelle/nook/region"
	"golang.org/x/text/language"
)

const (
	// doubutsuNoMori is the common reference for Doubutsu no Mori.
	doubutsuNoMori = "DoubutsuNoMori"
)

var (
	// doubutsuNoMoriNameAmericanEnglish represents Doubutsu no Mori's name in American English.
	doubutsuNoMoriNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Doubutsu no Mori",
	}
)

var (
	// doubutsuNoMoriReleaseDateJapan represents Doubutsu no Mori's release date in Japan.
	doubutsuNoMoriReleaseDateJapan = nook.ReleaseDate{
		Day:    14,
		Month:  time.April,
		Region: region.Japan,
		Year:   2001,
	}
)

var (
	// doubutsuNoMoriName contains the localized names of Doubutsu no Mori.
	doubutsuNoMoriName = nook.Languages{
		language.AmericanEnglish: doubutsuNoMoriNameAmericanEnglish,
	}
)

var (
	// doubutsuNoMoriPlatforms contains Doubutsu no Mori's release platforms in deterministic key order.
	doubutsuNoMoriPlatforms = []nook.Platform{
		platform.Nintendo64,
	}
)

var (
	// doubutsuNoMoriReleaseDates contains Doubutsu no Mori's regional release history in chronological order.
	doubutsuNoMoriReleaseDates = []nook.ReleaseDate{
		doubutsuNoMoriReleaseDateJapan,
	}
)

var (
	// DoubutsuNoMori represents Doubutsu no Mori.
	DoubutsuNoMori = nook.Game{
		Key:          nook.Key(doubutsuNoMori),
		Name:         doubutsuNoMoriName,
		Platforms:    doubutsuNoMoriPlatforms,
		ReleaseDates: doubutsuNoMoriReleaseDates,
		ReleaseOrder: 1,
	}
)
