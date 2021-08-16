package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	bearcubNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bearcub"}
)

var (
	bearcubName = nook.Languages{
		language.AmericanEnglish: bearcubNameAmericanEnglish}
)

var (
	Bearcub = nook.Animal{
		Name: bearcubName}
)
