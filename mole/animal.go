package mole

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	moleNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Mole"}
)

var (
	moleName = nook.Languages{
		language.AmericanEnglish: moleNameAmericanEnglish}
)

var (
	Mole = nook.Animal{
		Name: moleName}
)
