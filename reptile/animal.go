package reptile

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	reptileNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Reptile"}
)

var (
	reptileName = nook.Languages{
		language.AmericanEnglish: reptileNameAmericanEnglish}
)

var (
	Reptile = nook.Animal{
		Name: reptileName}
)
