package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // frillneckedlizard is the common reference for the Frillneckedlizard animal type.
    frillneckedlizard = "Frillneckedlizard"
)

var (
	// frillneckedlizardNameAmericanEnglish represents the American English name for the Frillneckedlizard animal type.
	frillneckedlizardNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    frillneckedlizard,
	}
)

var (
	// frillneckedlizardName contains the names of the Frillneckedlizard animal type in different languages.
	frillneckedlizardName = nook.Languages{
		language.AmericanEnglish: frillneckedlizardNameAmericanEnglish,
	}
)

var (
	// Frillneckedlizard represents an frillneckedlizard animal in the Animal Crossing series.
	Frillneckedlizard = nook.Animal{
		Key:  nook.Key(frillneckedlizard),
		Name: frillneckedlizardName,
	}
)
