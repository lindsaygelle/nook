package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	tortoiseNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Tortoise"}
)

var (
	tortoiseName = nook.Languages{
		language.AmericanEnglish: tortoiseNameAmericanEnglish}
)

var (
	Tortoise = nook.Animal{
		Name: tortoiseName}
)
