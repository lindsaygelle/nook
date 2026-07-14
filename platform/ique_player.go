package platform

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// iQuePlayer is the common reference for the iQue Player platform.
	iQuePlayer = "IQuePlayer"
)

var (
	// iQuePlayerNameAmericanEnglish represents the iQue Player platform's name in American English.
	iQuePlayerNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "iQue Player",
	}
)

var (
	// iQuePlayerName contains the localized names of the iQue Player platform.
	iQuePlayerName = nook.Languages{
		language.AmericanEnglish: iQuePlayerNameAmericanEnglish,
	}
)

var (
	// IQuePlayer represents the iQue Player platform in the Animal Crossing series.
	IQuePlayer = nook.Platform{
		Key:  nook.Key(iQuePlayer),
		Name: iQuePlayerName,
	}
)
