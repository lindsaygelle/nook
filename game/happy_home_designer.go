package game

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// happyHomeDesigner is the common reference for Animal Crossing: Happy Home Designer.
	happyHomeDesigner = "HappyHomeDesigner"
)

var (
	// happyHomeDesignerNameAmericanEnglish represents Animal Crossing: Happy Home Designer's name in American English.
	happyHomeDesignerNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Animal Crossing: Happy Home Designer",
	}
)

var (
	// happyHomeDesignerName contains the localized names of Animal Crossing: Happy Home Designer.
	happyHomeDesignerName = nook.Languages{
		language.AmericanEnglish: happyHomeDesignerNameAmericanEnglish,
	}
)

var (
	// HappyHomeDesigner represents Animal Crossing: Happy Home Designer.
	HappyHomeDesigner = nook.Game{
		Key:  nook.Key(happyHomeDesigner),
		Name: happyHomeDesignerName,
	}
)
