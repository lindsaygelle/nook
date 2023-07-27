package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // owl is the common reference for the Owl animal type.
    owl = "Owl"
)

var (
	// owlNameAmericanEnglish represents the American English name for the Owl animal type.
	owlNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    owl,
	}
)

var (
	// owlName contains the names of the Owl animal type in different languages.
	owlName = nook.Languages{
		language.AmericanEnglish: owlNameAmericanEnglish,
	}
)

var (
	// Owl represents an owl animal in the Animal Crossing series.
	Owl = nook.Animal{
		Key:  nook.Key(owl),
		Name: owlName,
	}
)
