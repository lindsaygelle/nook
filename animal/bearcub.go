package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	bearcubNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bear cub"}
)

var (
	bearcubName = nook.Languages{
		language.AmericanEnglish: bearcubNameAmericanEnglish}
)

var (
	Bearcub = nook.Animal{
		Key:  nook.Key("Bearcub"),
		Name: bearcubName}
)
