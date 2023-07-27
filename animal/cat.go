package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // cat is the common reference for the Cat animal type.
    cat = "Cat"
)

var (
	// catNameAmericanEnglish represents the American English name for the Cat animal type.
	catNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    cat,
	}
)

var (
	// catName contains the names of the Cat animal type in different languages.
	catName = nook.Languages{
		language.AmericanEnglish: catNameAmericanEnglish,
	}
)

var (
	// Cat represents an cat animal in the Animal Crossing series.
	Cat = nook.Animal{
		Key:  nook.Key(cat),
		Name: catName,
	}
)
