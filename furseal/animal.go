package furseal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	fursealNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Furseal"}
)

var (
	fursealName = nook.Languages{
		language.AmericanEnglish: fursealNameAmericanEnglish}
)

var (
	Furseal = nook.Animal{
		Name: fursealName}
)
