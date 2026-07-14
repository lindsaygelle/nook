package platform

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// android is the common reference for the Android platform.
	android = "Android"
)

var (
	// androidNameAmericanEnglish represents the Android platform's name in American English.
	androidNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Android",
	}
)

var (
	// androidName contains the localized names of the Android platform.
	androidName = nook.Languages{
		language.AmericanEnglish: androidNameAmericanEnglish,
	}
)

var (
	// Android represents the Android platform in the Animal Crossing series.
	Android = nook.Platform{
		Key:  nook.Key(android),
		Name: androidName,
	}
)
