package platform

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// nintendo3DS is the common reference for the Nintendo 3DS platform.
	nintendo3DS = "Nintendo3DS"
)

var (
	// nintendo3DSNameAmericanEnglish represents the Nintendo 3DS platform's name in American English.
	nintendo3DSNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Nintendo 3DS",
	}
)

var (
	// nintendo3DSName contains the localized names of the Nintendo 3DS platform.
	nintendo3DSName = nook.Languages{
		language.AmericanEnglish: nintendo3DSNameAmericanEnglish,
	}
)

var (
	// Nintendo3DS represents the Nintendo 3DS platform in the Animal Crossing series.
	Nintendo3DS = nook.Platform{
		Key:  nook.Key(nintendo3DS),
		Name: nintendo3DSName,
	}
)
