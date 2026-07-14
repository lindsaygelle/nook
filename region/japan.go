package region

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// japan is the common reference for the Japan region.
	japan = "Japan"
)

var (
	// japanNameAmericanEnglish represents the Japan region's name in American English.
	japanNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Japan",
	}
)

var (
	// japanName contains the localized names of the Japan region.
	japanName = nook.Languages{
		language.AmericanEnglish: japanNameAmericanEnglish,
	}
)

var (
	// Japan represents the Japan region in the Animal Crossing series.
	Japan = nook.Region{
		Key:  nook.Key(japan),
		Name: japanName,
	}
)
