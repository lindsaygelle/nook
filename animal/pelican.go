package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // pelican is the common reference for the Pelican animal type.
    pelican = "Pelican"
)

var (
	// pelicanNameAmericanEnglish represents the American English name for the Pelican animal type.
	pelicanNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    pelican,
	}
)

var (
	// pelicanName contains the names of the Pelican animal type in different languages.
	pelicanName = nook.Languages{
		language.AmericanEnglish: pelicanNameAmericanEnglish,
	}
)

var (
	// Pelican represents an pelican animal in the Animal Crossing series.
	Pelican = nook.Animal{
		Key:  nook.Key(pelican),
		Name: pelicanName,
	}
)
