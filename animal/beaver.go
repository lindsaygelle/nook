package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	beaverNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Beaver"}
)

var (
	beaverName = nook.Languages{
		language.AmericanEnglish: beaverNameAmericanEnglish}
)

var (
	Beaver = nook.Animal{
		Name: beaverName}
)
