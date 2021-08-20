package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	rabbitNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Rabbit"}
)

var (
	rabbitName = nook.Languages{
		language.AmericanEnglish: rabbitNameAmericanEnglish}
)

var (
	Rabbit = nook.Animal{
		Key:  nook.Key("Rabbit"),
		Name: rabbitName}
)
