package platform

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// ios is the common reference for the iOS platform.
	ios = "IOS"
)

var (
	// iosNameAmericanEnglish represents the iOS platform's name in American English.
	iosNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "iOS",
	}
)

var (
	// iosName contains the localized names of the iOS platform.
	iosName = nook.Languages{
		language.AmericanEnglish: iosNameAmericanEnglish,
	}
)

var (
	// IOS represents the iOS platform in the Animal Crossing series.
	IOS = nook.Platform{
		Key:  nook.Key(ios),
		Name: iosName,
	}
)
