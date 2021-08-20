package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	pigeonNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Pigeon"}
)

var (
	pigeonName = nook.Languages{
		language.AmericanEnglish: pigeonNameAmericanEnglish}
)

var (
	Pigeon = nook.Animal{
		Key:  nook.Key("Pigeon"),
		Name: pigeonName}
)
