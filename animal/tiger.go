package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // tiger is the common reference for the Tiger animal type.
    tiger = "Tiger"
)

var (
	// tigerNameAmericanEnglish represents the American English name for the Tiger animal type.
	tigerNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    tiger,
	}
)

var (
	// tigerName contains the names of the Tiger animal type in different languages.
	tigerName = nook.Languages{
		language.AmericanEnglish: tigerNameAmericanEnglish,
	}
)

var (
	// Tiger represents an tiger animal in the Animal Crossing series.
	Tiger = nook.Animal{
		Key:  nook.Key(tiger),
		Name: tigerName,
	}
)
