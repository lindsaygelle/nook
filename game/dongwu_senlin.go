package game

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/platform"
	"github.com/lindsaygelle/nook/region"
	"golang.org/x/text/language"
)

const (
	// dongwuSenlin is the common reference for Dongwu Senlin.
	dongwuSenlin = "DongwuSenlin"
)

var (
	// dongwuSenlinNameAmericanEnglish represents Dongwu Senlin's name in American English.
	dongwuSenlinNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Dongwu Senlin",
	}
)

var (
	// dongwuSenlinReleaseDateChina represents Dongwu Senlin's release date in China.
	dongwuSenlinReleaseDateChina = nook.ReleaseDate{
		Day:    1,
		Month:  time.June,
		Region: region.China,
		Year:   2006,
	}
)

var (
	// dongwuSenlinName contains the localized names of Dongwu Senlin.
	dongwuSenlinName = nook.Languages{
		language.AmericanEnglish: dongwuSenlinNameAmericanEnglish,
	}
)

var (
	// dongwuSenlinPlatforms contains Dongwu Senlin's release platforms in deterministic key order.
	dongwuSenlinPlatforms = []nook.Platform{
		platform.IQuePlayer,
	}
)

var (
	// dongwuSenlinReleaseDates contains Dongwu Senlin's regional release history in chronological order.
	dongwuSenlinReleaseDates = []nook.ReleaseDate{
		dongwuSenlinReleaseDateChina,
	}
)

var (
	// DongwuSenlin represents Dongwu Senlin.
	DongwuSenlin = nook.Game{
		Key:          nook.Key(dongwuSenlin),
		Name:         dongwuSenlinName,
		Platforms:    dongwuSenlinPlatforms,
		ReleaseDates: dongwuSenlinReleaseDates,
		ReleaseOrder: 5,
	}
)
