package role

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// islander is the common reference for the Islander special-character role.
	islander = "Islander"
)

var (
	// islanderNameAmericanEnglish represents the Islander role's name in American English.
	islanderNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Islander",
	}
)

var (
	// islanderName contains the localized names of the Islander role.
	islanderName = nook.Languages{
		language.AmericanEnglish: islanderNameAmericanEnglish,
	}
)

var (
	// Islander represents the Islander special-character role.
	Islander = nook.Role{
		Key:  nook.Key(islander),
		Name: islanderName,
	}
)
