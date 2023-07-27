package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // pig is the common reference for the Pig animal type.
    pig = "Pig"
)

var (
	// pigNameAmericanEnglish represents the American English name for the Pig animal type.
	pigNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    pig,
	}
)

var (
	// pigName contains the names of the Pig animal type in different languages.
	pigName = nook.Languages{
		language.AmericanEnglish: pigNameAmericanEnglish,
	}
)

var (
	// Pig represents an pig animal in the Animal Crossing series.
	Pig = nook.Animal{
		Key:  nook.Key(pig),
		Name: pigName,
	}
)
