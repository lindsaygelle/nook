package role

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// government is the common reference for the Government special-character role.
	government = "Government"
)

var (
	// governmentNameAmericanEnglish represents the Government role's name in American English.
	governmentNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Government",
	}
)

var (
	// governmentName contains the localized names of the Government role.
	governmentName = nook.Languages{
		language.AmericanEnglish: governmentNameAmericanEnglish,
	}
)

var (
	// Government represents the Government special-character role.
	Government = nook.Role{
		Key:  nook.Key(government),
		Name: governmentName,
	}
)
