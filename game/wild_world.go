package game

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// wildWorld is the common reference for Animal Crossing: Wild World.
	wildWorld = "WildWorld"
)

var (
	// wildWorldNameAmericanEnglish represents Animal Crossing: Wild World's name in American English.
	wildWorldNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Animal Crossing: Wild World",
	}
)

var (
	// wildWorldName contains the localized names of Animal Crossing: Wild World.
	wildWorldName = nook.Languages{
		language.AmericanEnglish: wildWorldNameAmericanEnglish,
	}
)

var (
	// WildWorld represents Animal Crossing: Wild World.
	WildWorld = nook.Game{
		Key:  nook.Key(wildWorld),
		Name: wildWorldName,
	}
)
