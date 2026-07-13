package game

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// doubutsuNoMoriEPlus is the common reference for Doubutsu no Mori e+.
	doubutsuNoMoriEPlus = "DoubutsuNoMoriEPlus"
)

var (
	// doubutsuNoMoriEPlusNameAmericanEnglish represents Doubutsu no Mori e+'s name in American English.
	doubutsuNoMoriEPlusNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Doubutsu no Mori e+",
	}
)

var (
	// doubutsuNoMoriEPlusName contains the localized names of Doubutsu no Mori e+.
	doubutsuNoMoriEPlusName = nook.Languages{
		language.AmericanEnglish: doubutsuNoMoriEPlusNameAmericanEnglish,
	}
)

var (
	// DoubutsuNoMoriEPlus represents Doubutsu no Mori e+.
	DoubutsuNoMoriEPlus = nook.Game{
		Key:  nook.Key(doubutsuNoMoriEPlus),
		Name: doubutsuNoMoriEPlusName,
	}
)
