package goat

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	iggyBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	iggyCode = nook.Code{
		Value: ""}
)

var (
	iggyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Iggy"}

	iggyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	iggyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	iggyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Cabri"}

	iggyGermanName = nook.Name{
		Language: language.German,
		Value:    "Georg"}

	iggyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Caprice"}

	iggyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "オトマツ"}

	iggyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	iggyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	iggyRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	iggySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "炯炯"}

	iggySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Montés"}

	iggyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
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
	iggyCharacter = nook.Character{
		Animal:   Goat,
		Birthday: iggyBirthday,
		Code:     iggyCode,
		Gender:   nook.Male,
		Name:     iggyName}
)

var (
	iggyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "paaally"}

	iggyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	iggyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	iggyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mon jojo"}

	iggyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "kumpel"}

	iggyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "salameee"}

	iggyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ですじゃ"}

	iggyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	iggyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	iggyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	iggySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "Unknown"}

	iggySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "coleguiita"}

	iggyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	iggyPhrase = nook.Languages{
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
	Iggy = nook.Villager{
		Character:   iggyCharacter,
		Personality: nook.Jock,
		Phrase:      iggyPhrase}
)
