package goat

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	sherbBirthday = nook.Birthday{
		Day:   18,
		Month: time.January}
)

var (
	sherbCode = nook.Code{
		Value: "goa09"}
)

var (
	sherbAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Sherb"}

	sherbCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Capribââââille"}

	sherbDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Sherboefff"}

	sherbFrenchName = nook.Name{
		Language: language.French,
		Value:    "Capribââââille"}

	sherbGermanName = nook.Name{
		Language: language.German,
		Value:    "Morpheusschlummer"}

	sherbItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Capraldobeeewn"}

	sherbJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "レムふわぁ"}

	sherbKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "래미흐아앙"}

	sherbLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Morfeobostezzz"}

	sherbRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Шербмемеша"}

	sherbSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "雷姆轻飘"}

	sherbSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Morfeobostezzz"}

	sherbTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "雷姆輕飄"}
)

var (
	sherbName = nook.Languages{
		language.AmericanEnglish:      sherbAmericanEnglishName,
		language.CanadianFrench:       sherbCanadianFrenchName,
		language.Dutch:                sherbDutchName,
		language.French:               sherbFrenchName,
		language.German:               sherbGermanName,
		language.Italian:              sherbItalianName,
		language.Japanese:             sherbJapaneseName,
		language.Korean:               sherbKoreanName,
		language.LatinAmericanSpanish: sherbLatinAmericanSpanishName,
		language.Russian:              sherbRussianName,
		language.SimplifiedChinese:    sherbSimplifiedChineseName,
		language.Spanish:              sherbSpanishName,
		language.TraditionalChinese:   sherbTraditionalChineseName}
)

var (
	sherbCharacter = nook.Character{
		Animal:   Goat,
		Birthday: sherbBirthday,
		Code:     sherbCode,
		Gender:   nook.Male,
		Name:     sherbName}
)

var (
	sherbAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "bawwww"}

	sherbCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "bââââille"}

	sherbDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "oefff"}

	sherbFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "bââââille"}

	sherbGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "schlummer"}

	sherbItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "beeewn"}

	sherbJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ふわぁ"}

	sherbKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "흐아앙"}

	sherbLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "bostezzz"}

	sherbRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "мемеша"}

	sherbSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "轻飘"}

	sherbSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "bostezzz"}

	sherbTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "輕飄"}
)

var (
	sherbPhrase = nook.Languages{
		language.AmericanEnglish:      sherbAmericanEnglishName,
		language.CanadianFrench:       sherbCanadianFrenchName,
		language.Dutch:                sherbDutchName,
		language.French:               sherbFrenchName,
		language.German:               sherbGermanName,
		language.Italian:              sherbItalianName,
		language.Japanese:             sherbJapaneseName,
		language.Korean:               sherbKoreanName,
		language.LatinAmericanSpanish: sherbLatinAmericanSpanishName,
		language.Russian:              sherbRussianName,
		language.SimplifiedChinese:    sherbSimplifiedChineseName,
		language.Spanish:              sherbSpanishName,
		language.TraditionalChinese:   sherbTraditionalChineseName}
)

var (
	Sherb = nook.Villager{
		Character:   sherbCharacter,
		Personality: nook.Lazy,
		Phrase:      sherbPhrase}
)
