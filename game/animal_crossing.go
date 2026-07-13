package game

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// animalCrossing is the common reference for Animal Crossing.
	animalCrossing = "AnimalCrossing"
)

var (
	// animalCrossingNameAmericanEnglish represents Animal Crossing's name in American English.
	animalCrossingNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Animal Crossing",
	}
)

var (
	// animalCrossingName contains the localized names of Animal Crossing.
	animalCrossingName = nook.Languages{
		language.AmericanEnglish: animalCrossingNameAmericanEnglish,
	}
)

var (
	// AnimalCrossing represents Animal Crossing.
	AnimalCrossing = nook.Game{
		Key:  nook.Key(animalCrossing),
		Name: animalCrossingName,
	}
)
