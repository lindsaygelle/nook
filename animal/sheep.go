package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	sheepNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Sheep"}
)

var (
	sheepName = nook.Languages{
		language.AmericanEnglish: sheepNameAmericanEnglish}
)

var (
	Sheep = nook.Animal{
		Key:  nook.Key("Sheep"),
		Name: sheepName}
)
