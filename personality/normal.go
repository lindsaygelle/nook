package personality

import (
	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

const (
	// normal is the common reference for the Normal personality.
	normal = "Normal"
)

var (
	// normalAmericanEnglishName represents the American English name for the Normal personality.
	normalAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Normal"}
)

var (
	// normalName contains the names of the Normal personality in different languages.
	normalName = nook.Languages{
		language.AmericanEnglish: normalAmericanEnglishName}
)

var (
	// Normal represents the Normal personality in the Animal Crossing series.
	Normal = nook.Personality{
		Key:  nook.Key(normal),
		Name: normalName}
)
