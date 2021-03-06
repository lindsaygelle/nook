package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	ostrichNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Ostrich"}
)

var (
	ostrichName = nook.Languages{
		language.AmericanEnglish: ostrichNameAmericanEnglish}
)

var (
	Ostrich = nook.Animal{
		Key:  nook.Key("Ostrich"),
		Name: ostrichName}
)
