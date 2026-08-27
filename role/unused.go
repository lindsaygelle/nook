package role

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// unused is the common reference for the Unused special-character role.
	unused = "Unused"
)

var (
	// unusedNameAmericanEnglish represents the Unused role's name in American English.
	unusedNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Unused",
	}
)

var (
	// unusedName contains the localized names of the Unused role.
	unusedName = nook.Languages{
		language.AmericanEnglish: unusedNameAmericanEnglish,
	}
)

var (
	// Unused represents the Unused special-character role.
	Unused = nook.Role{
		Key:  nook.Key(unused),
		Name: unusedName,
	}
)
