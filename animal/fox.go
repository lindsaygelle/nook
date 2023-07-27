package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // fox is the common reference for the Fox animal type.
    fox = "Fox"
)

var (
	// foxNameAmericanEnglish represents the American English name for the Fox animal type.
	foxNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    fox,
	}
)

var (
	// foxName contains the names of the Fox animal type in different languages.
	foxName = nook.Languages{
		language.AmericanEnglish: foxNameAmericanEnglish,
	}
)

var (
	// Fox represents an fox animal in the Animal Crossing series.
	Fox = nook.Animal{
		Key:  nook.Key(fox),
		Name: foxName,
	}
)
