package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	hamsterNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Hamster"}
)

var (
	hamsterName = nook.Languages{
		language.AmericanEnglish: hamsterNameAmericanEnglish}
)

var (
	Hamster = nook.Animal{
		Key:  nook.Key("Hamster"),
		Name: hamsterName}
)
