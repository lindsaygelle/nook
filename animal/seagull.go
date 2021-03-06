package animal

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	seagullNameAmericanEnglish = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Seagull"}
)

var (
	seagullName = nook.Languages{
		language.AmericanEnglish: seagullNameAmericanEnglish}
)

var (
	Seagull = nook.Animal{
		Key:  nook.Key("Seagull"),
		Name: seagullName}
)
