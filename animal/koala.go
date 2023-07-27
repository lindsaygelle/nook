package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // koala is the common reference for the Koala animal type.
    koala = "Koala"
)

var (
	// koalaNameAmericanEnglish represents the American English name for the Koala animal type.
	koalaNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    koala,
	}
)

var (
	// koalaName contains the names of the Koala animal type in different languages.
	koalaName = nook.Languages{
		language.AmericanEnglish: koalaNameAmericanEnglish,
	}
)

var (
	// Koala represents an koala animal in the Animal Crossing series.
	Koala = nook.Animal{
		Key:  nook.Key(koala),
		Name: koalaName,
	}
)
