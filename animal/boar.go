package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // boar is the common reference for the Boar animal type.
    boar = "Boar"
)

var (
	// boarNameAmericanEnglish represents the American English name for the Boar animal type.
	boarNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    boar,
	}
)

var (
	// boarName contains the names of the Boar animal type in different languages.
	boarName = nook.Languages{
		language.AmericanEnglish: boarNameAmericanEnglish,
	}
)

var (
	// Boar represents an boar animal in the Animal Crossing series.
	Boar = nook.Animal{
		Key:  nook.Key(boar),
		Name: boarName,
	}
)
