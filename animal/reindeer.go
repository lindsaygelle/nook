package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	reindeerNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Reindeer"}
)

var (
	reindeerName = nook.Languages{
		language.AmericanEnglish: reindeerNameAmericanEnglish}
)

var (
	Reindeer = nook.Animal{
		Key:  nook.Key("Reindeer"),
		Name: reindeerName}
)
