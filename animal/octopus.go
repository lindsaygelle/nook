package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // octopus is the common reference for the Octopus animal type.
    octopus = "Octopus"
)

var (
	// octopusNameAmericanEnglish represents the American English name for the Octopus animal type.
	octopusNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    octopus,
	}
)

var (
	// octopusName contains the names of the Octopus animal type in different languages.
	octopusName = nook.Languages{
		language.AmericanEnglish: octopusNameAmericanEnglish,
	}
)

var (
	// Octopus represents an octopus animal in the Animal Crossing series.
	Octopus = nook.Animal{
		Key:  nook.Key(octopus),
		Name: octopusName,
	}
)
