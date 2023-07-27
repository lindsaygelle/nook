package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // axolotl is the common reference for the Axolotl animal type.
    axolotl = "Axolotl"
)

var (
	// axolotlNameAmericanEnglish represents the American English name for the Axolotl animal type.
	axolotlNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    axolotl,
	}
)

var (
	// axolotlName contains the names of the Axolotl animal type in different languages.
	axolotlName = nook.Languages{
		language.AmericanEnglish: axolotlNameAmericanEnglish,
	}
)

var (
	// Axolotl represents an axolotl animal in the Animal Crossing series.
	Axolotl = nook.Animal{
		Key:  nook.Key(axolotl),
		Name: axolotlName,
	}
)
