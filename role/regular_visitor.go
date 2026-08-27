package role

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// regularVisitor is the common reference for the RegularVisitor special-character role.
	regularVisitor = "RegularVisitor"
)

var (
	// regularVisitorNameAmericanEnglish represents the RegularVisitor role's name in American English.
	regularVisitorNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Regular Visitor",
	}
)

var (
	// regularVisitorName contains the localized names of the RegularVisitor role.
	regularVisitorName = nook.Languages{
		language.AmericanEnglish: regularVisitorNameAmericanEnglish,
	}
)

var (
	// RegularVisitor represents the RegularVisitor special-character role.
	RegularVisitor = nook.Role{
		Key:  nook.Key(regularVisitor),
		Name: regularVisitorName,
	}
)
