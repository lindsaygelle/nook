package koala

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "Koalinemon trésor"}

	sydneyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Sydneyzonnetje"}

	sydneyFrenchName = nook.Name{
		Language: language.French,
		Value:    "Koalinemon trésor"}

	sydneyGermanName = nook.Name{
		Language: language.German,
		Value:    "Silkesternchen"}

	sydneyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Sidneycalipto"}

	sydneyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "シドニーだコアラ"}

	sydneyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "시드니발그레"}

	sydneyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Sidneycalipto"}

	sydneyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Сиднисолнышко"}

	sydneySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "思颖无尾熊"}

	sydneySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Sidneysúper"}

	sydneyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "思穎無尾熊"}
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
		Animal:   Koala,
		Birthday: sydneyBirthday,
		Code:     sydneyCode,
		Gender:   nook.Female,
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
		Personality: nook.Normal,
		Phrase:      sydneyPhrase}
)
