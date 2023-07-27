package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // bear is the common reference for the Bear animal type.
    bear = "Bear"
)

var (
	// bearNameAmericanEnglish represents the American English name for the Bear animal type.
	bearNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    bear,
	}
)

var (
	// bearName contains the names of the Bear animal type in different languages.
	bearName = nook.Languages{
		language.AmericanEnglish: bearNameAmericanEnglish,
	}
)

var (
	// Bear represents an bear animal in the Animal Crossing series.
	Bear = nook.Animal{
		Key:  nook.Key(bear),
		Name: bearName,
	}
)
