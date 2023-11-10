package alligator

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

// alfonsoBirthday represents Alfonso's birthday (June 9th).
var alfonsoBirthday = nook.Birthday{
	Day:   9,
	Month: time.June,
}

// alfonsoCode represents Alfonso's unique code ("crd00").
var alfonsoCode = nook.Code{
	Value: "crd00",
}

// alfonsoName represents Alfonso's name in different languages.
var alfonsoName = nook.Languages{
	language.AmericanEnglish:      {Language: language.AmericanEnglish, Value: "Alfonso"},
	language.CanadianFrench:       {Language: language.CanadianFrench, Value: "Alphonse"},
	language.Dutch:                {Language: language.Dutch, Value: "Alfonso"},
	language.French:               {Language: language.French, Value: "Alphonse"},
	language.German:               {Language: language.German, Value: "Markus"},
	language.Italian:              {Language: language.Italian, Value: "Alfonso"},
	language.Japanese:             {Language: language.Japanese, Value: "アルベルト"},
	language.Korean:               {Language: language.Korean, Value: "알베르트"},
	language.LatinAmericanSpanish: {Language: language.LatinAmericanSpanish, Value: "Kaimán"},
	language.Russian:              {Language: language.Russian, Value: "Альфонсо"},
	language.SimplifiedChinese:    {Language: language.SimplifiedChinese, Value: "阿泥"},
	language.Spanish:              {Language: language.Spanish, Value: "Kaimán"},
	language.TraditionalChinese:   {Language: language.TraditionalChinese, Value: "阿泥"},
}

// alfonsoPhrase represents Alfonso's phrases in different languages.
var alfonsoPhrase = nook.Languages{
	language.AmericanEnglish:      {Language: language.AmericanEnglish, Value: "it's a me"},
	language.CanadianFrench:       {Language: language.CanadianFrench, Value: "moi moi"},
	language.Dutch:                {Language: language.Dutch, Value: "it'sa me"},
	language.French:               {Language: language.French, Value: "moi moi"},
	language.German:               {Language: language.German, Value: "ahhhrg"},
	language.Italian:              {Language: language.Italian, Value: "proprio"},
	language.Japanese:             {Language: language.Japanese, Value: "だワニ"},
	language.Korean:               {Language: language.Korean, Value: "나아거"},
	language.LatinAmericanSpanish: {Language: language.LatinAmericanSpanish, Value: "ñam-ñam"},
	language.Russian:              {Language: language.Russian, Value: "это я"},
	language.SimplifiedChinese:    {Language: language.SimplifiedChinese, Value: "鳄泥"},
	language.Spanish:              {Language: language.Spanish, Value: "ñam-ñam"},
	language.TraditionalChinese:   {Language: language.TraditionalChinese, Value: "鱷泥"},
}

// alfonsoCharacter represents Alfonso's character information.
var alfonsoCharacter = nook.Character{
	Animal:   animal.Alligator,
	Birthday: alfonsoBirthday,
	Code:     alfonsoCode,
	Key:      character.Alfonso,
	Gender:   gender.Male,
	Name:     alfonsoName,
	Special:  false,
}

// Alfonso represents the character Alfonso with his complete information.
var Alfonso = nook.Villager{
	Character:   alfonsoCharacter,
	Personality: personality.Lazy,
	Phrase:      alfonsoPhrase,
}
