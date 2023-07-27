package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // chameleon is the common reference for the Chameleon animal type.
    chameleon = "Chameleon"
)

var (
	// chameleonNameAmericanEnglish represents the American English name for the Chameleon animal type.
	chameleonNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    chameleon,
	}
)

var (
	// chameleonName contains the names of the Chameleon animal type in different languages.
	chameleonName = nook.Languages{
		language.AmericanEnglish: chameleonNameAmericanEnglish,
	}
)

var (
	// Chameleon represents an chameleon animal in the Animal Crossing series.
	Chameleon = nook.Animal{
		Key:  nook.Key(chameleon),
		Name: chameleonName,
	}
)
