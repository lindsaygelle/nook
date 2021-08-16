package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	tapirNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Tapir"}
)

var (
	tapirName = nook.Languages{
		language.AmericanEnglish: tapirNameAmericanEnglish}
)

var (
	Tapir = nook.Animal{
		Name: tapirName}
)
