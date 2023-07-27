package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // alpaca is the common reference for the Alpaca animal type.
    alpaca = "Alpaca"
)

var (
	// alpacaNameAmericanEnglish represents the American English name for the Alpaca animal type.
	alpacaNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    alpaca,
	}
)

var (
	// alpacaName contains the names of the Alpaca animal type in different languages.
	alpacaName = nook.Languages{
		language.AmericanEnglish: alpacaNameAmericanEnglish,
	}
)

var (
	// Alpaca represents an alpaca animal in the Animal Crossing series.
	Alpaca = nook.Animal{
		Key:  nook.Key(alpaca),
		Name: alpacaName,
	}
)
