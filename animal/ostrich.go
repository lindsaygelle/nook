package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // ostrich is the common reference for the Ostrich animal type.
    ostrich = "Ostrich"
)

var (
	// ostrichNameAmericanEnglish represents the American English name for the Ostrich animal type.
	ostrichNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ostrich,
	}
)

var (
	// ostrichName contains the names of the Ostrich animal type in different languages.
	ostrichName = nook.Languages{
		language.AmericanEnglish: ostrichNameAmericanEnglish,
	}
)

var (
	// Ostrich represents an ostrich animal in the Animal Crossing series.
	Ostrich = nook.Animal{
		Key:  nook.Key(ostrich),
		Name: ostrichName,
	}
)
