package koala

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
	sydneyBirthday = nook.Birthday{
		Day:   21,
		Month: time.June}
)

var (
	sydneyCode = nook.Code{
		Value: "kal03"}
)

var (
	sydneyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Sydney"}

	sydneyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Koaline"}

	sydneyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Sydney"}

	sydneyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Koaline"}

	sydneyGermanName = nook.Name{
		Language: language.German,
		Value:    "Silke"}

	sydneyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Sidney"}

	sydneyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "シドニー"}

	sydneyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "시드니"}

	sydneyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Sidney"}

	sydneyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Сидни"}

	sydneySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "思颖"}

	sydneySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Sidney"}

	sydneyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "思穎"}
)

var (
	sydneyName = nook.Languages{
		language.AmericanEnglish:      sydneyAmericanEnglishName,
		language.CanadianFrench:       sydneyCanadianFrenchName,
		language.Dutch:                sydneyDutchName,
		language.French:               sydneyFrenchName,
		language.German:               sydneyGermanName,
		language.Italian:              sydneyItalianName,
		language.Japanese:             sydneyJapaneseName,
		language.Korean:               sydneyKoreanName,
		language.LatinAmericanSpanish: sydneyLatinAmericanSpanishName,
		language.Russian:              sydneyRussianName,
		language.SimplifiedChinese:    sydneySimplifiedChineseName,
		language.Spanish:              sydneySpanishName,
		language.TraditionalChinese:   sydneyTraditionalChineseName}
)

var (
	sydneyCharacter = nook.Character{
		Animal:   animal.Koala,
		Birthday: sydneyBirthday,
		Code:     sydneyCode,
		Key:      character.Sydney,
		Gender:   gender.Female,
		Name:     sydneyName}
)

var (
	sydneyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "sunshine"}

	sydneyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "mon trésor"}

	sydneyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "zonnetje"}

	sydneyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "mon trésor"}

	sydneyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "sternchen"}

	sydneyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "calipto"}

	sydneyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "だコアラ"}

	sydneyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "발그레"}

	sydneyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "calipto"}

	sydneyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "солнышко"}

	sydneySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "无尾熊"}

	sydneySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "súper"}

	sydneyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "無尾熊"}
)

var (
	sydneyPhrase = nook.Languages{
		language.AmericanEnglish:      sydneyAmericanEnglishName,
		language.CanadianFrench:       sydneyCanadianFrenchName,
		language.Dutch:                sydneyDutchName,
		language.French:               sydneyFrenchName,
		language.German:               sydneyGermanName,
		language.Italian:              sydneyItalianName,
		language.Japanese:             sydneyJapaneseName,
		language.Korean:               sydneyKoreanName,
		language.LatinAmericanSpanish: sydneyLatinAmericanSpanishName,
		language.Russian:              sydneyRussianName,
		language.SimplifiedChinese:    sydneySimplifiedChineseName,
		language.Spanish:              sydneySpanishName,
		language.TraditionalChinese:   sydneyTraditionalChineseName}
)

var (
	Sydney = nook.Villager{
		Character:   sydneyCharacter,
		Personality: personality.Normal,
		Phrase:      sydneyPhrase}
)
