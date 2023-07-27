package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // horse is the common reference for the Horse animal type.
    horse = "Horse"
)

var (
	// horseNameAmericanEnglish represents the American English name for the Horse animal type.
	horseNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    horse,
	}
)

var (
	// horseName contains the names of the Horse animal type in different languages.
	horseName = nook.Languages{
		language.AmericanEnglish: horseNameAmericanEnglish,
	}
)

var (
	// Horse represents an horse animal in the Animal Crossing series.
	Horse = nook.Animal{
		Key:  nook.Key(horse),
		Name: horseName,
	}
)
