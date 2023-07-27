package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // pigeon is the common reference for the Pigeon animal type.
    pigeon = "Pigeon"
)

var (
	// pigeonNameAmericanEnglish represents the American English name for the Pigeon animal type.
	pigeonNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    pigeon,
	}
)

var (
	// pigeonName contains the names of the Pigeon animal type in different languages.
	pigeonName = nook.Languages{
		language.AmericanEnglish: pigeonNameAmericanEnglish,
	}
)

var (
	// Pigeon represents an pigeon animal in the Animal Crossing series.
	Pigeon = nook.Animal{
		Key:  nook.Key(pigeon),
		Name: pigeonName,
	}
)
