package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // giraffe is the common reference for the Giraffe animal type.
    giraffe = "Giraffe"
)

var (
	// giraffeNameAmericanEnglish represents the American English name for the Giraffe animal type.
	giraffeNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    giraffe,
	}
)

var (
	// giraffeName contains the names of the Giraffe animal type in different languages.
	giraffeName = nook.Languages{
		language.AmericanEnglish: giraffeNameAmericanEnglish,
	}
)

var (
	// Giraffe represents an giraffe animal in the Animal Crossing series.
	Giraffe = nook.Animal{
		Key:  nook.Key(giraffe),
		Name: giraffeName,
	}
)
