package game

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// doubutsuNoMoriPlus is the common reference for Doubutsu no Mori+.
	doubutsuNoMoriPlus = "DoubutsuNoMoriPlus"
)

var (
	// doubutsuNoMoriPlusNameAmericanEnglish represents Doubutsu no Mori+'s name in American English.
	doubutsuNoMoriPlusNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Doubutsu no Mori+",
	}
)

var (
	// doubutsuNoMoriPlusName contains the localized names of Doubutsu no Mori+.
	doubutsuNoMoriPlusName = nook.Languages{
		language.AmericanEnglish: doubutsuNoMoriPlusNameAmericanEnglish,
	}
)

var (
	// DoubutsuNoMoriPlus represents Doubutsu no Mori+.
	DoubutsuNoMoriPlus = nook.Game{
		Key:  nook.Key(doubutsuNoMoriPlus),
		Name: doubutsuNoMoriPlusName,
	}
)
