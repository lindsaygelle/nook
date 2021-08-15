package frog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	sunnyBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	sunnyCode = nook.Code{
		Value: ""}
)

var (
	sunnyAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Sunny"}

	sunnyCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	sunnyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	sunnyFrenchName = nook.Name{
		Language: language.French,
		Value:    "N/A"}

	sunnyGermanName = nook.Name{
		Language: language.German,
		Value:    "N/A"}

	sunnyItalianName = nook.Name{
		Language: language.Italian,
		Value:    "N/A"}

	sunnyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "サニー"}

	sunnyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	sunnyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	sunnyRussianName = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	sunnySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	sunnySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "N/A"}

	sunnyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	sunnyName = nook.Languages{
		language.AmericanEnglish:      sunnyAmericanEnglishName,
		language.CanadianFrench:       sunnyCanadianFrenchName,
		language.Dutch:                sunnyDutchName,
		language.French:               sunnyFrenchName,
		language.German:               sunnyGermanName,
		language.Italian:              sunnyItalianName,
		language.Japanese:             sunnyJapaneseName,
		language.Korean:               sunnyKoreanName,
		language.LatinAmericanSpanish: sunnyLatinAmericanSpanishName,
		language.Russian:              sunnyRussianName,
		language.SimplifiedChinese:    sunnySimplifiedChineseName,
		language.Spanish:              sunnySpanishName,
		language.TraditionalChinese:   sunnyTraditionalChineseName}
)

var (
	sunnyCharacter = nook.Character{
		Animal:   Frog,
		Birthday: sunnyBirthday,
		Code:     sunnyCode,
		Gender:   nook.Female,
		Name:     sunnyName}
)

var (
	sunnyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	sunnyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	sunnyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	sunnyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "N/A"}

	sunnyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "N/A"}

	sunnyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "N/A"}

	sunnyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でちょ"}

	sunnyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	sunnyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	sunnyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	sunnySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	sunnySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "N/A"}

	sunnyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
)

var (
	sunnyPhrase = nook.Languages{
		language.AmericanEnglish:      sunnyAmericanEnglishName,
		language.CanadianFrench:       sunnyCanadianFrenchName,
		language.Dutch:                sunnyDutchName,
		language.French:               sunnyFrenchName,
		language.German:               sunnyGermanName,
		language.Italian:              sunnyItalianName,
		language.Japanese:             sunnyJapaneseName,
		language.Korean:               sunnyKoreanName,
		language.LatinAmericanSpanish: sunnyLatinAmericanSpanishName,
		language.Russian:              sunnyRussianName,
		language.SimplifiedChinese:    sunnySimplifiedChineseName,
		language.Spanish:              sunnySpanishName,
		language.TraditionalChinese:   sunnyTraditionalChineseName}
)

var (
	Sunny = nook.Villager{
		Character:   sunnyCharacter,
		Personality: nook.Normal,
		Phrase:      sunnyPhrase}
)
