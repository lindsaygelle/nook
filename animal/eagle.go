package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // eagle is the common reference for the Eagle animal type.
    eagle = "Eagle"
)

var (
	// eagleNameAmericanEnglish represents the American English name for the Eagle animal type.
	eagleNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    eagle,
	}
)

var (
	// eagleName contains the names of the Eagle animal type in different languages.
	eagleName = nook.Languages{
		language.AmericanEnglish: eagleNameAmericanEnglish,
	}
)

var (
	// Eagle represents an eagle animal in the Animal Crossing series.
	Eagle = nook.Animal{
		Key:  nook.Key(eagle),
		Name: eagleName,
	}
)
