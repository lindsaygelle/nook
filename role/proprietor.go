package role

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// proprietor is the common reference for the Proprietor special-character role.
	proprietor = "Proprietor"
)

var (
	// proprietorNameAmericanEnglish represents the Proprietor role's name in American English.
	proprietorNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Proprietor",
	}
)

var (
	// proprietorName contains the localized names of the Proprietor role.
	proprietorName = nook.Languages{
		language.AmericanEnglish: proprietorNameAmericanEnglish,
	}
)

var (
	// Proprietor represents the Proprietor special-character role.
	Proprietor = nook.Role{
		Key:  nook.Key(proprietor),
		Name: proprietorName,
	}
)
