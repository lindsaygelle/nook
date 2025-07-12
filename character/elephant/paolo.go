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
	// paoloBirthday represents paolo birthday.
	paoloBirthday = nook.Birthday{
		Day:   5,
		Month: time.May}
)

var (
	// paoloCode represents paolo code.
	paoloCode = nook.Code{
		Value: "elp05"}
)

var (
	// paoloAmericanEnglishName represents paolo american english name.
	paoloAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Paolo"}

	// paoloCanadianFrenchName represents paolo canadian french name.
	paoloCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Paolo"}

	// paoloDutchName represents paolo dutch name.
	paoloDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Paolo"}

	// paoloFrenchName represents paolo french name.
	paoloFrenchName = nook.Name{
		Language: language.French,
		Value:    "Paolo"}

	// paoloGermanName represents paolo german name.
	paoloGermanName = nook.Name{
		Language: language.German,
		Value:    "Paolo"}

	// paoloItalianName represents paolo italian name.
	paoloItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Paolino"}

	// paoloJapaneseName represents paolo japanese name.
	paoloJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "パオロ"}

	// paoloKoreanName represents paolo korean name.
	paoloKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "파올로"}

	// paoloLatinAmericanSpanishName represents paolo latin american spanish name.
	paoloLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Paolo"}

	// paoloRussianName represents paolo russian name.
	paoloRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Паоло"}

	// paoloSimplifiedChineseName represents paolo simplified chinese name.
	paoloSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "保罗"}

	// paoloSpanishName represents paolo spanish name.
	paoloSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Paolo"}

	// paoloTraditionalChineseName represents paolo traditional chinese name.
	paoloTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "保羅"}
)

var (
	// paoloName represents paolo name.
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
	// paoloCharacter represents paolo character.
	paoloCharacter = nook.Character{
		Animal:   animal.Elephant,
		Birthday: paoloBirthday,
		Code:     paoloCode,
		Key:      character.Paolo,
		Gender:   gender.Male,
		Name:     paoloName,
		Special:  false}
)

var (
	// paoloAmericanEnglishPhrase represents paolo american english phrase.
	paoloAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "pal"}

	// paoloCanadianFrenchPhrase represents paolo canadian french phrase.
	paoloCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "fanfan"}

	// paoloDutchPhrase represents paolo dutch phrase.
	paoloDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "gabber"}

	// paoloFrenchPhrase represents paolo french phrase.
	paoloFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "fanfan"}

	// paoloGermanPhrase represents paolo german phrase.
	paoloGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "pruuust"}

	// paoloItalianPhrase represents paolo italian phrase.
	paoloItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "pistaaaa"}

	// paoloJapanesePhrase represents paolo japanese phrase.
	paoloJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "パオ"}

	// paoloKoreanPhrase represents paolo korean phrase.
	paoloKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "파오"}

	// paoloLatinAmericanSpanishPhrase represents paolo latin american spanish phrase.
	paoloLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "pruuuf"}

	// paoloRussianPhrase represents paolo russian phrase.
	paoloRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "друг"}

	// paoloSimplifiedChinesePhrase represents paolo simplified chinese phrase.
	paoloSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "保罗"}

	// paoloSpanishPhrase represents paolo spanish phrase.
	paoloSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "amigante"}

	// paoloTraditionalChinesePhrase represents paolo traditional chinese phrase.
	paoloTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "保羅"}
)

var (
	// paoloPhrase represents paolo phrase.
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
	// Paolo represents paolo.
	Paolo = nook.Villager{
		Character:   paoloCharacter,
		Personality: personality.Lazy,
		Phrase:      paoloPhrase}
)
