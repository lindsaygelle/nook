package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
    // doc is the common reference for the Doc animal type.
    doc = "Doc"
)

var (
	// docNameAmericanEnglish represents the American English name for the Doc animal type.
	docNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    doc,
	}
)

var (
	// docName contains the names of the Doc animal type in different languages.
	docName = nook.Languages{
		language.AmericanEnglish: docNameAmericanEnglish,
	}
)

var (
	// Doc represents an doc animal in the Animal Crossing series.
	Doc = nook.Animal{
		Key:  nook.Key(doc),
		Name: docName,
	}
)
