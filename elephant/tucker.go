package elephant

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	tuckerBirthday = nook.Birthday{
		Day:   7,
		Month: time.September}
)

var (
	tuckerCode = nook.Code{
		Value: "elp09"}
)

var (
	tuckerAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Tucker"}

	tuckerCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Barrymammouth"}

	tuckerDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Tuckerpluisie"}

	tuckerFrenchName = nook.Name{
		Language: language.French,
		Value:    "Barrymammouth"}

	tuckerGermanName = nook.Name{
		Language: language.German,
		Value:    "Thorstenpuuuuuh"}

	tuckerItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Sventolobarrir"}

	tuckerJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "はじめもじゃ"}

	tuckerKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "맘모뿌뿌"}

	tuckerLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Pirolopurupú"}

	tuckerRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Такербивень"}

	tuckerSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "阿原毛毛"}

	tuckerSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Pirolopurupú"}

	tuckerTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "阿原毛毛"}
)

var (
	tuckerName = nook.Languages{
		language.AmericanEnglish:      tuckerAmericanEnglishName,
		language.CanadianFrench:       tuckerCanadianFrenchName,
		language.Dutch:                tuckerDutchName,
		language.French:               tuckerFrenchName,
		language.German:               tuckerGermanName,
		language.Italian:              tuckerItalianName,
		language.Japanese:             tuckerJapaneseName,
		language.Korean:               tuckerKoreanName,
		language.LatinAmericanSpanish: tuckerLatinAmericanSpanishName,
		language.Russian:              tuckerRussianName,
		language.SimplifiedChinese:    tuckerSimplifiedChineseName,
		language.Spanish:              tuckerSpanishName,
		language.TraditionalChinese:   tuckerTraditionalChineseName}
)

var (
	tuckerCharacter = nook.Character{
		Animal:   Elephant,
		Birthday: tuckerBirthday,
		Code:     tuckerCode,
		Gender:   nook.Male,
		Name:     tuckerName}
)

var (
	tuckerAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "fuzzers"}

	tuckerCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mammouth"}

	tuckerDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "pluisie"}

	tuckerFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mammouth"}

	tuckerGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "puuuuuh"}

	tuckerItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "barrir"}

	tuckerJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "もじゃ"}

	tuckerKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "뿌뿌"}

	tuckerLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "purupú"}

	tuckerRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "бивень"}

	tuckerSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "毛毛"}

	tuckerSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "purupú"}

	tuckerTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "毛毛"}
)

var (
	tuckerPhrase = nook.Languages{
		language.AmericanEnglish:      tuckerAmericanEnglishName,
		language.CanadianFrench:       tuckerCanadianFrenchName,
		language.Dutch:                tuckerDutchName,
		language.French:               tuckerFrenchName,
		language.German:               tuckerGermanName,
		language.Italian:              tuckerItalianName,
		language.Japanese:             tuckerJapaneseName,
		language.Korean:               tuckerKoreanName,
		language.LatinAmericanSpanish: tuckerLatinAmericanSpanishName,
		language.Russian:              tuckerRussianName,
		language.SimplifiedChinese:    tuckerSimplifiedChineseName,
		language.Spanish:              tuckerSpanishName,
		language.TraditionalChinese:   tuckerTraditionalChineseName}
)

var (
	Tucker = nook.Villager{
		Character:   tuckerCharacter,
		Personality: nook.Lazy,
		Phrase:      tuckerPhrase}
)
