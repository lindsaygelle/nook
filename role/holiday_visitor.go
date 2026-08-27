package role

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// holidayVisitor is the common reference for the HolidayVisitor special-character role.
	holidayVisitor = "HolidayVisitor"
)

var (
	// holidayVisitorNameAmericanEnglish represents the HolidayVisitor role's name in American English.
	holidayVisitorNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Holiday Visitor",
	}
)

var (
	// holidayVisitorName contains the localized names of the HolidayVisitor role.
	holidayVisitorName = nook.Languages{
		language.AmericanEnglish: holidayVisitorNameAmericanEnglish,
	}
)

var (
	// HolidayVisitor represents the HolidayVisitor special-character role.
	HolidayVisitor = nook.Role{
		Key:  nook.Key(holidayVisitor),
		Name: holidayVisitorName,
	}
)
