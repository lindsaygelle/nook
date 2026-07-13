package game

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// cityFolk is the common reference for Animal Crossing: City Folk.
	cityFolk = "CityFolk"
)

var (
	// cityFolkNameAmericanEnglish represents Animal Crossing: City Folk's name in American English.
	cityFolkNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Animal Crossing: City Folk",
	}
)

var (
	// cityFolkName contains the localized names of Animal Crossing: City Folk.
	cityFolkName = nook.Languages{
		language.AmericanEnglish: cityFolkNameAmericanEnglish,
	}
)

var (
	// CityFolk represents Animal Crossing: City Folk.
	CityFolk = nook.Game{
		Key:  nook.Key(cityFolk),
		Name: cityFolkName,
	}
)
