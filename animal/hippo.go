package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // hippo is the common reference for the Hippo animal type.
    hippo = "Hippo"
)

var (
	// hippoNameAmericanEnglish represents the American English name for the Hippo animal type.
	hippoNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    hippo,
	}
)

var (
	// hippoName contains the names of the Hippo animal type in different languages.
	hippoName = nook.Languages{
		language.AmericanEnglish: hippoNameAmericanEnglish,
	}
)

var (
	// Hippo represents an hippo animal in the Animal Crossing series.
	Hippo = nook.Animal{
		Key:  nook.Key(hippo),
		Name: hippoName,
	}
)
