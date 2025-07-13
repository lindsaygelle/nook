package goat

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
	// iggyBirthday represents iggy birthday.
	iggyBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	// iggyCode represents iggy code.
	iggyCode = nook.Code{
		Value: ""}
)

var (
	// iggyAmericanEnglishName represents iggy american english name.
	iggyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Iggy"}

	// iggyCanadianFrenchName represents iggy canadian french name.
	iggyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// iggyDutchName represents iggy dutch name.
	iggyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// iggyFrenchName represents iggy french name.
	iggyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Cabri"}

	// iggyGermanName represents iggy german name.
	iggyGermanName = nook.Name{
		Language: language.German,
		Value:    "Georg"}

	// iggyItalianName represents iggy italian name.
	iggyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Caprice"}

	// iggyJapaneseName represents iggy japanese name.
	iggyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "オトマツ"}

	// iggyKoreanName represents iggy korean name.
	iggyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// iggyLatinAmericanSpanishName represents iggy latin american spanish name.
	iggyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// iggyRussianName represents iggy russian name.
	iggyRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// iggySimplifiedChineseName represents iggy simplified chinese name.
	iggySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "炯炯"}

	// iggySpanishName represents iggy spanish name.
	iggySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Montés"}

	// iggyTraditionalChineseName represents iggy traditional chinese name.
	iggyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// iggyName represents iggy name.
	iggyName = nook.Languages{
		language.AmericanEnglish:      iggyAmericanEnglishName,
		language.CanadianFrench:       iggyCanadianFrenchName,
		language.Dutch:                iggyDutchName,
		language.French:               iggyFrenchName,
		language.German:               iggyGermanName,
		language.Italian:              iggyItalianName,
		language.Japanese:             iggyJapaneseName,
		language.Korean:               iggyKoreanName,
		language.LatinAmericanSpanish: iggyLatinAmericanSpanishName,
		language.Russian:              iggyRussianName,
		language.SimplifiedChinese:    iggySimplifiedChineseName,
		language.Spanish:              iggySpanishName,
		language.TraditionalChinese:   iggyTraditionalChineseName}
)

var (
	// iggyCharacter represents iggy character.
	iggyCharacter = nook.Character{
		Animal:   animal.Goat,
		Birthday: iggyBirthday,
		Code:     iggyCode,
		Key:      character.Iggy,
		Gender:   gender.Male,
		Name:     iggyName,
		Special:  false}
)

var (
	// iggyAmericanEnglishPhrase represents iggy american english phrase.
	iggyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "paaally"}

	// iggyCanadianFrenchPhrase represents iggy canadian french phrase.
	iggyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	// iggyDutchPhrase represents iggy dutch phrase.
	iggyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	// iggyFrenchPhrase represents iggy french phrase.
	iggyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mon jojo"}

	// iggyGermanPhrase represents iggy german phrase.
	iggyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kumpel"}

	// iggyItalianPhrase represents iggy italian phrase.
	iggyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "salameee"}

	// iggyJapanesePhrase represents iggy japanese phrase.
	iggyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ですじゃ"}

	// iggyKoreanPhrase represents iggy korean phrase.
	iggyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	// iggyLatinAmericanSpanishPhrase represents iggy latin american spanish phrase.
	iggyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	// iggyRussianPhrase represents iggy russian phrase.
	iggyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	// iggySimplifiedChinesePhrase represents iggy simplified chinese phrase.
	iggySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Unknown"}

	// iggySpanishPhrase represents iggy spanish phrase.
	iggySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "coleguiita"}

	// iggyTraditionalChinesePhrase represents iggy traditional chinese phrase.
	iggyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	// iggyPhrase represents iggy phrase.
	iggyPhrase = nook.Languages{
		language.AmericanEnglish:      iggyAmericanEnglishPhrase,
		language.CanadianFrench:       iggyCanadianFrenchPhrase,
		language.Dutch:                iggyDutchPhrase,
		language.French:               iggyFrenchPhrase,
		language.German:               iggyGermanPhrase,
		language.Italian:              iggyItalianPhrase,
		language.Japanese:             iggyJapanesePhrase,
		language.Korean:               iggyKoreanPhrase,
		language.LatinAmericanSpanish: iggyLatinAmericanSpanishPhrase,
		language.Russian:              iggyRussianPhrase,
		language.SimplifiedChinese:    iggySimplifiedChinesePhrase,
		language.Spanish:              iggySpanishPhrase,
		language.TraditionalChinese:   iggyTraditionalChinesePhrase}
)

var (
	// Iggy represents iggy.
	Iggy = nook.Villager{
		Character:   iggyCharacter,
		Personality: personality.Jock,
		Phrase:      iggyPhrase}
)
