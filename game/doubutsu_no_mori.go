package game

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// doubutsuNoMori is the common reference for Doubutsu no Mori.
	doubutsuNoMori = "DoubutsuNoMori"
)

var (
	// doubutsuNoMoriNameAmericanEnglish represents Doubutsu no Mori's name in American English.
	doubutsuNoMoriNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Doubutsu no Mori",
	}
)

var (
	// doubutsuNoMoriName contains the localized names of Doubutsu no Mori.
	doubutsuNoMoriName = nook.Languages{
		language.AmericanEnglish: doubutsuNoMoriNameAmericanEnglish,
	}
)

var (
	// DoubutsuNoMori represents Doubutsu no Mori.
	DoubutsuNoMori = nook.Game{
		Key:  nook.Key(doubutsuNoMori),
		Name: doubutsuNoMoriName,
	}
)
