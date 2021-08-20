package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	birdNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bird"}
)

var (
	birdName = nook.Languages{
		language.AmericanEnglish: birdNameAmericanEnglish}
)

var (
	Bird = nook.Animal{
		Key:  nook.Key("Bird"),
		Name: birdName}
)
