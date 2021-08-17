package elephant

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
	paoloBirthday = nook.Birthday{
		Day:   5,
		Month: time.May}
)

var (
	paoloCode = nook.Code{
		Value: "elp05"}
)

var (
	paoloAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Paolo"}

	paoloCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Paolo"}

	paoloDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Paolo"}

	paoloFrenchName = nook.Name{
		Language: language.French,
		Value:    "Paolo"}

	paoloGermanName = nook.Name{
		Language: language.German,
		Value:    "Paolo"}

	paoloItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Paolino"}

	paoloJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "パオロ"}

	paoloKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "파올로"}

	paoloLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Paolo"}

	paoloRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Паоло"}

	paoloSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "保罗"}

	paoloSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Paolo"}

	paoloTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "保羅"}
)

var (
	paoloName = nook.Languages{
		language.AmericanEnglish:      paoloAmericanEnglishName,
		language.CanadianFrench:       paoloCanadianFrenchName,
		language.Dutch:                paoloDutchName,
		language.French:               paoloFrenchName,
		language.German:               paoloGermanName,
		language.Italian:              paoloItalianName,
		language.Japanese:             paoloJapaneseName,
		language.Korean:               paoloKoreanName,
		language.LatinAmericanSpanish: paoloLatinAmericanSpanishName,
		language.Russian:              paoloRussianName,
		language.SimplifiedChinese:    paoloSimplifiedChineseName,
		language.Spanish:              paoloSpanishName,
		language.TraditionalChinese:   paoloTraditionalChineseName}
)

var (
	paoloCharacter = nook.Character{
		Animal:   animal.Elephant,
		Birthday: paoloBirthday,
		Code:     paoloCode,
		Key:      character.Paolo,
		Gender:   gender.Male,
		Name:     paoloName}
)

var (
	paoloAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "pal"}

	paoloCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "fanfan"}

	paoloDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "gabber"}

	paoloFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "fanfan"}

	paoloGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "pruuust"}

	paoloItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "pistaaaa"}

	paoloJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "パオ"}

	paoloKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "파오"}

	paoloLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "pruuuf"}

	paoloRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "друг"}

	paoloSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "保罗"}

	paoloSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "amigante"}

	paoloTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "保羅"}
)

var (
	paoloPhrase = nook.Languages{
		language.AmericanEnglish:      paoloAmericanEnglishPhrase,
		language.CanadianFrench:       paoloCanadianFrenchPhrase,
		language.Dutch:                paoloDutchPhrase,
		language.French:               paoloFrenchPhrase,
		language.German:               paoloGermanPhrase,
		language.Italian:              paoloItalianPhrase,
		language.Japanese:             paoloJapanesePhrase,
		language.Korean:               paoloKoreanPhrase,
		language.LatinAmericanSpanish: paoloLatinAmericanSpanishPhrase,
		language.Russian:              paoloRussianPhrase,
		language.SimplifiedChinese:    paoloSimplifiedChinesePhrase,
		language.Spanish:              paoloSpanishPhrase,
		language.TraditionalChinese:   paoloTraditionalChinesePhrase}
)

var (
	Paolo = nook.Villager{
		Character:   paoloCharacter,
		Personality: personality.Lazy,
		Phrase:      paoloPhrase}
)
