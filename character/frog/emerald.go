package frog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/character"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	// emeraldBirthday represents emerald birthday.
	emeraldBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// emeraldCode represents emerald code.
	emeraldCode = nook.Code{
		Value: ""}
)

var (
	// emeraldAmericanEnglishName represents emerald american english name.
	emeraldAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Emerald"}

	// emeraldCanadianFrenchName represents emerald canadian french name.
	emeraldCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// emeraldDutchName represents emerald dutch name.
	emeraldDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// emeraldFrenchName represents emerald french name.
	emeraldFrenchName = nook.Name{
		Language: language.French,
		Value:    "Émeraude"}

	// emeraldGermanName represents emerald german name.
	emeraldGermanName = nook.Name{
		Language: language.German,
		Value:    "Emma"}

	// emeraldItalianName represents emerald italian name.
	emeraldItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Smeralda"}

	// emeraldJapaneseName represents emerald japanese name.
	emeraldJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ケロミ"}

	// emeraldKoreanName represents emerald korean name.
	emeraldKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// emeraldLatinAmericanSpanishName represents emerald latin american spanish name.
	emeraldLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// emeraldRussianName represents emerald russian name.
	emeraldRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// emeraldSimplifiedChineseName represents emerald simplified chinese name.
	emeraldSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "呱呱"}

	// emeraldSpanishName represents emerald spanish name.
	emeraldSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Espe"}

	// emeraldTraditionalChineseName represents emerald traditional chinese name.
	emeraldTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// emeraldName represents emerald name.
	emeraldName = nook.Languages{
		language.AmericanEnglish:      emeraldAmericanEnglishName,
		language.CanadianFrench:       emeraldCanadianFrenchName,
		language.Dutch:                emeraldDutchName,
		language.French:               emeraldFrenchName,
		language.German:               emeraldGermanName,
		language.Italian:              emeraldItalianName,
		language.Japanese:             emeraldJapaneseName,
		language.Korean:               emeraldKoreanName,
		language.LatinAmericanSpanish: emeraldLatinAmericanSpanishName,
		language.Russian:              emeraldRussianName,
		language.SimplifiedChinese:    emeraldSimplifiedChineseName,
		language.Spanish:              emeraldSpanishName,
		language.TraditionalChinese:   emeraldTraditionalChineseName}
)

var (
	// emeraldCharacter represents emerald character.
	emeraldCharacter = nook.Character{
		Animal:   animal.Frog,
		Birthday: emeraldBirthday,
		Code:     emeraldCode,
		Key:      character.Emerald,
		Gender:   gender.Female,
		Name:     emeraldName,
		Special:  false}
)

var (
	// emeraldAmericanEnglishPhrase represents emerald american english phrase.
	emeraldAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "sproing"}

	// emeraldCanadianFrenchPhrase represents emerald canadian french phrase.
	emeraldCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// emeraldDutchPhrase represents emerald dutch phrase.
	emeraldDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// emeraldFrenchPhrase represents emerald french phrase.
	emeraldFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "sproiiing"}

	// emeraldGermanPhrase represents emerald german phrase.
	emeraldGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "boioioing"}

	// emeraldItalianPhrase represents emerald italian phrase.
	emeraldItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "boing"}

	// emeraldJapanesePhrase represents emerald japanese phrase.
	emeraldJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だケロ"}

	// emeraldKoreanPhrase represents emerald korean phrase.
	emeraldKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// emeraldLatinAmericanSpanishPhrase represents emerald latin american spanish phrase.
	emeraldLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// emeraldRussianPhrase represents emerald russian phrase.
	emeraldRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// emeraldSimplifiedChinesePhrase represents emerald simplified chinese phrase.
	emeraldSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "咕咕"}

	// emeraldSpanishPhrase represents emerald spanish phrase.
	emeraldSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "esproing"}

	// emeraldTraditionalChinesePhrase represents emerald traditional chinese phrase.
	emeraldTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// emeraldPhrase represents emerald phrase.
	emeraldPhrase = nook.Languages{
		language.AmericanEnglish:      emeraldAmericanEnglishPhrase,
		language.CanadianFrench:       emeraldCanadianFrenchPhrase,
		language.Dutch:                emeraldDutchPhrase,
		language.French:               emeraldFrenchPhrase,
		language.German:               emeraldGermanPhrase,
		language.Italian:              emeraldItalianPhrase,
		language.Japanese:             emeraldJapanesePhrase,
		language.Korean:               emeraldKoreanPhrase,
		language.LatinAmericanSpanish: emeraldLatinAmericanSpanishPhrase,
		language.Russian:              emeraldRussianPhrase,
		language.SimplifiedChinese:    emeraldSimplifiedChinesePhrase,
		language.Spanish:              emeraldSpanishPhrase,
		language.TraditionalChinese:   emeraldTraditionalChinesePhrase}
)

var (
	// Emerald represents emerald.
	Emerald = nook.Villager{
		Character:   emeraldCharacter,
		Personality: personality.Normal,
		Phrase:      emeraldPhrase}
)
