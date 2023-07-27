package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // penguin is the common reference for the Penguin animal type.
    penguin = "Penguin"
)

var (
	// penguinNameAmericanEnglish represents the American English name for the Penguin animal type.
	penguinNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    penguin,
	}
)

var (
	// penguinName contains the names of the Penguin animal type in different languages.
	penguinName = nook.Languages{
		language.AmericanEnglish: penguinNameAmericanEnglish,
	}
)

var (
	// Penguin represents an penguin animal in the Animal Crossing series.
	Penguin = nook.Animal{
		Key:  nook.Key(penguin),
		Name: penguinName,
	}
)
