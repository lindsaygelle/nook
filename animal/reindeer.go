package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // reindeer is the common reference for the Reindeer animal type.
    reindeer = "Reindeer"
)

var (
	// reindeerNameAmericanEnglish represents the American English name for the Reindeer animal type.
	reindeerNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    reindeer,
	}
)

var (
	// reindeerName contains the names of the Reindeer animal type in different languages.
	reindeerName = nook.Languages{
		language.AmericanEnglish: reindeerNameAmericanEnglish,
	}
)

var (
	// Reindeer represents an reindeer animal in the Animal Crossing series.
	Reindeer = nook.Animal{
		Key:  nook.Key(reindeer),
		Name: reindeerName,
	}
)
