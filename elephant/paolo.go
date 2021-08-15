package elephant

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Paolofanfan"}

	paoloDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Paologabber"}

	paoloFrenchName = nook.Name{
		Language: language.French,
		Value:    "Paolofanfan"}

	paoloGermanName = nook.Name{
		Language: language.German,
		Value:    "Paolopruuust"}

	paoloItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Paolinopistaaaa"}

	paoloJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "パオロパオ"}

	paoloKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "파올로파오"}

	paoloLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Paolopruuuf"}

	paoloRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Паолодруг"}

	paoloSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "保罗保罗"}

	paoloSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Paoloamigante"}

	paoloTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "保羅保羅"}
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
		Animal:   Elephant,
		Birthday: paoloBirthday,
		Code:     paoloCode,
		Gender:   nook.Male,
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
	Paolo = nook.Villager{
		Character:   paoloCharacter,
		Personality: nook.Lazy,
		Phrase:      paoloPhrase}
)
