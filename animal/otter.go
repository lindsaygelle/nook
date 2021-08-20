package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	otterNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Otter"}
)

var (
	otterName = nook.Languages{
		language.AmericanEnglish: otterNameAmericanEnglish}
)

var (
	Otter = nook.Animal{
		Key:  nook.Key("Otter"),
		Name: otterName}
)
