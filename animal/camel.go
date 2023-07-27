package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // camel is the common reference for the Camel animal type.
    camel = "Camel"
)

var (
	// camelNameAmericanEnglish represents the American English name for the Camel animal type.
	camelNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    camel,
	}
)

var (
	// camelName contains the names of the Camel animal type in different languages.
	camelName = nook.Languages{
		language.AmericanEnglish: camelNameAmericanEnglish,
	}
)

var (
	// Camel represents an camel animal in the Animal Crossing series.
	Camel = nook.Animal{
		Key:  nook.Key(camel),
		Name: camelName,
	}
)
