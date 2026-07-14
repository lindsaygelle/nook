package platform

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// nintendo64 is the common reference for the Nintendo 64 platform.
	nintendo64 = "Nintendo64"
)

var (
	// nintendo64NameAmericanEnglish represents the Nintendo 64 platform's name in American English.
	nintendo64NameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Nintendo 64",
	}
)

var (
	// nintendo64Name contains the localized names of the Nintendo 64 platform.
	nintendo64Name = nook.Languages{
		language.AmericanEnglish: nintendo64NameAmericanEnglish,
	}
)

var (
	// Nintendo64 represents the Nintendo 64 platform in the Animal Crossing series.
	Nintendo64 = nook.Platform{
		Key:  nook.Key(nintendo64),
		Name: nintendo64Name,
	}
)
