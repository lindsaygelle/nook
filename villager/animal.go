package villager

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	villagerNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Villager"}
)

var (
	villagerName = nook.Languages{
		language.AmericanEnglish: villagerNameAmericanEnglish}
)

var (
	Villager = nook.Animal{
		Name: villagerName}
)
