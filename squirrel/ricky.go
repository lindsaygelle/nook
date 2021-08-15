package squirrel

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	rickyBirthday = nook.Birthday{
		Day:   14,
		Month: time.September}
)

var (
	rickyCode = nook.Code{
		Value: "squ10"}
)

var (
	rickyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Ricky"}

	rickyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Rockydingo"}

	rickyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Rickymafnoot"}

	rickyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Rockydingo"}

	rickyGermanName = nook.Name{
		Language: language.German,
		Value:    "Ronnyknorz"}

	rickyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Sberlonocciola"}

	rickyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "カジロウキッ"}

	rickyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "갈가리사각사각"}

	rickyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Altramuzcrac-crac"}

	rickyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Рикиарахис"}

	rickySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "嘉治炯炯"}

	rickySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Altramuzcáscaras"}

	rickyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "嘉治炯炯"}
)

var (
	rickyName = nook.Languages{
		language.AmericanEnglish:      rickyAmericanEnglishName,
		language.CanadianFrench:       rickyCanadianFrenchName,
		language.Dutch:                rickyDutchName,
		language.French:               rickyFrenchName,
		language.German:               rickyGermanName,
		language.Italian:              rickyItalianName,
		language.Japanese:             rickyJapaneseName,
		language.Korean:               rickyKoreanName,
		language.LatinAmericanSpanish: rickyLatinAmericanSpanishName,
		language.Russian:              rickyRussianName,
		language.SimplifiedChinese:    rickySimplifiedChineseName,
		language.Spanish:              rickySpanishName,
		language.TraditionalChinese:   rickyTraditionalChineseName}
)

var (
	rickyCharacter = nook.Character{
		Animal:   Squirrel,
		Birthday: rickyBirthday,
		Code:     rickyCode,
		Gender:   nook.Male,
		Name:     rickyName}
)

var (
	rickyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "nutcase"}

	rickyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "dingo"}

	rickyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "mafnoot"}

	rickyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "dingo"}

	rickyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "knorz"}

	rickyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "nocciola"}

	rickyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "キッ"}

	rickyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "사각사각"}

	rickyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "crac-crac"}

	rickyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "арахис"}

	rickySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "炯炯"}

	rickySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "cáscaras"}

	rickyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "炯炯"}
)

var (
	rickyPhrase = nook.Languages{
		language.AmericanEnglish:      rickyAmericanEnglishName,
		language.CanadianFrench:       rickyCanadianFrenchName,
		language.Dutch:                rickyDutchName,
		language.French:               rickyFrenchName,
		language.German:               rickyGermanName,
		language.Italian:              rickyItalianName,
		language.Japanese:             rickyJapaneseName,
		language.Korean:               rickyKoreanName,
		language.LatinAmericanSpanish: rickyLatinAmericanSpanishName,
		language.Russian:              rickyRussianName,
		language.SimplifiedChinese:    rickySimplifiedChineseName,
		language.Spanish:              rickySpanishName,
		language.TraditionalChinese:   rickyTraditionalChineseName}
)

var (
	Ricky = nook.Villager{
		Character:   rickyCharacter,
		Personality: nook.Cranky,
		Phrase:      rickyPhrase}
)
