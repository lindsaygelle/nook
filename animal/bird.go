package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // bird is the common reference for the Bird animal type.
    bird = "Bird"
)

var (
	// birdNameAmericanEnglish represents the American English name for the Bird animal type.
	birdNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    bird,
	}
)

var (
	// birdName contains the names of the Bird animal type in different languages.
	birdName = nook.Languages{
		language.AmericanEnglish: birdNameAmericanEnglish,
	}
)

var (
	// Bird represents an bird animal in the Animal Crossing series.
	Bird = nook.Animal{
		Key:  nook.Key(bird),
		Name: birdName,
	}
)
