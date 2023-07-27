package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // bearcub is the common reference for the Bearcub animal type.
    bearcub = "Bearcub"
)

var (
	// bearcubNameAmericanEnglish represents the American English name for the Bearcub animal type.
	bearcubNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    bearcub,
	}
)

var (
	// bearcubName contains the names of the Bearcub animal type in different languages.
	bearcubName = nook.Languages{
		language.AmericanEnglish: bearcubNameAmericanEnglish,
	}
)

var (
	// Bearcub represents an bearcub animal in the Animal Crossing series.
	Bearcub = nook.Animal{
		Key:  nook.Key(bearcub),
		Name: bearcubName,
	}
)
