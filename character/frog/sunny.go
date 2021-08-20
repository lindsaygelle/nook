package frog

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
		Value:    ""}

	sunnyDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	sunnyFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	sunnyGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	sunnyItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	sunnyJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "サニー"}

	sunnyKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	sunnyLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	sunnyRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	sunnySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	sunnySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	sunnyTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
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
		Animal:   animal.Frog,
		Birthday: sunnyBirthday,
		Code:     sunnyCode,
		Key:      character.Sunny,
		Gender:   gender.Female,
		Name:     sunnyName,
		Special:  false}
)

var (
	sunnyAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "でちょ"}

	sunnyCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	sunnyDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	sunnyFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	sunnyGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	sunnyItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	sunnyJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "でちょ"}

	sunnyKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	sunnyLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	sunnyRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	sunnySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	sunnySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	sunnyTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	sunnyPhrase = nook.Languages{
		language.AmericanEnglish:      sunnyAmericanEnglishPhrase,
		language.CanadianFrench:       sunnyCanadianFrenchPhrase,
		language.Dutch:                sunnyDutchPhrase,
		language.French:               sunnyFrenchPhrase,
		language.German:               sunnyGermanPhrase,
		language.Italian:              sunnyItalianPhrase,
		language.Japanese:             sunnyJapanesePhrase,
		language.Korean:               sunnyKoreanPhrase,
		language.LatinAmericanSpanish: sunnyLatinAmericanSpanishPhrase,
		language.Russian:              sunnyRussianPhrase,
		language.SimplifiedChinese:    sunnySimplifiedChinesePhrase,
		language.Spanish:              sunnySpanishPhrase,
		language.TraditionalChinese:   sunnyTraditionalChinesePhrase}
)

var (
	Sunny = nook.Villager{
		Character:   sunnyCharacter,
		Personality: personality.Normal,
		Phrase:      sunnyPhrase}
)
