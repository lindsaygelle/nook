package gamecategory

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// mobile is the common reference for the Mobile game category.
	mobile = "Mobile"
)

var (
	// mobileNameAmericanEnglish represents the Mobile game category's name in American English.
	mobileNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Mobile",
	}
)

var (
	// mobileName contains the localized names of the Mobile game category.
	mobileName = nook.Languages{
		language.AmericanEnglish: mobileNameAmericanEnglish,
	}
)

var (
	// Mobile represents the Mobile game category in the Animal Crossing series.
	Mobile = nook.GameCategory{
		Key:  nook.Key(mobile),
		Name: mobileName,
	}
)
