package game

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

func newGame(key, name string) nook.Game {
	return nook.Game{
		Key: nook.Key(key),
		Name: nook.Languages{
			language.AmericanEnglish: {
				Language: language.AmericanEnglish,
				Value:    name,
			},
		},
	}
}

var (
	// AmiiboFestival represents Animal Crossing: amiibo Festival.
	AmiiboFestival = newGame("AmiiboFestival", "Animal Crossing: amiibo Festival")

	// AnimalCrossing represents Animal Crossing.
	AnimalCrossing = newGame("AnimalCrossing", "Animal Crossing")

	// CityFolk represents Animal Crossing: City Folk.
	CityFolk = newGame("CityFolk", "Animal Crossing: City Folk")

	// DongwuSenlin represents Dongwu Senlin.
	DongwuSenlin = newGame("DongwuSenlin", "Dongwu Senlin")

	// DoubutsuNoMori represents Doubutsu no Mori.
	DoubutsuNoMori = newGame("DoubutsuNoMori", "Doubutsu no Mori")

	// DoubutsuNoMoriEPlus represents Doubutsu no Mori e+.
	DoubutsuNoMoriEPlus = newGame("DoubutsuNoMoriEPlus", "Doubutsu no Mori e+")

	// DoubutsuNoMoriPlus represents Doubutsu no Mori+.
	DoubutsuNoMoriPlus = newGame("DoubutsuNoMoriPlus", "Doubutsu no Mori+")

	// HappyHomeDesigner represents Animal Crossing: Happy Home Designer.
	HappyHomeDesigner = newGame("HappyHomeDesigner", "Animal Crossing: Happy Home Designer")

	// NewHorizons represents Animal Crossing: New Horizons.
	NewHorizons = newGame("NewHorizons", "Animal Crossing: New Horizons")

	// NewLeaf represents Animal Crossing: New Leaf.
	NewLeaf = newGame("NewLeaf", "Animal Crossing: New Leaf")

	// PocketCamp represents Animal Crossing: Pocket Camp.
	PocketCamp = newGame("PocketCamp", "Animal Crossing: Pocket Camp")

	// WildWorld represents Animal Crossing: Wild World.
	WildWorld = newGame("WildWorld", "Animal Crossing: Wild World")
)
