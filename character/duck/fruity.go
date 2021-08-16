package duck

import (
	"time"

	"github.com/lindsaygelle/nook"
	"github.com/lindsaygelle/nook/animal"
	"github.com/lindsaygelle/nook/gender"
	"github.com/lindsaygelle/nook/personality"
	"golang.org/x/text/language"
)

var (
	fruityBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	fruityCode = nook.Code{
		Value: ""}
)

var (
	fruityAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Fruity"}

	fruityCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	fruityDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	fruityFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	fruityGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	fruityItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	fruityJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "フルーティー"}

	fruityKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	fruityLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	fruityRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	fruitySimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	fruitySpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	fruityTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	fruityName = nook.Languages{
		language.AmericanEnglish:      fruityAmericanEnglishName,
		language.CanadianFrench:       fruityCanadianFrenchName,
		language.Dutch:                fruityDutchName,
		language.French:               fruityFrenchName,
		language.German:               fruityGermanName,
		language.Italian:              fruityItalianName,
		language.Japanese:             fruityJapaneseName,
		language.Korean:               fruityKoreanName,
		language.LatinAmericanSpanish: fruityLatinAmericanSpanishName,
		language.Russian:              fruityRussianName,
		language.SimplifiedChinese:    fruitySimplifiedChineseName,
		language.Spanish:              fruitySpanishName,
		language.TraditionalChinese:   fruityTraditionalChineseName}
)

var (
	fruityCharacter = nook.Character{
		Animal:   animal.Duck,
		Birthday: fruityBirthday,
		Code:     fruityCode,
		Gender:   gender.Male,
		Name:     fruityName}
)

var (
	fruityAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "グァバ"}

	fruityCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	fruityDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	fruityFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	fruityGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	fruityItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	fruityJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "グァバ"}

	fruityKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	fruityLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	fruityRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	fruitySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	fruitySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	fruityTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	fruityPhrase = nook.Languages{
		language.AmericanEnglish:      fruityAmericanEnglishName,
		language.CanadianFrench:       fruityCanadianFrenchName,
		language.Dutch:                fruityDutchName,
		language.French:               fruityFrenchName,
		language.German:               fruityGermanName,
		language.Italian:              fruityItalianName,
		language.Japanese:             fruityJapaneseName,
		language.Korean:               fruityKoreanName,
		language.LatinAmericanSpanish: fruityLatinAmericanSpanishName,
		language.Russian:              fruityRussianName,
		language.SimplifiedChinese:    fruitySimplifiedChineseName,
		language.Spanish:              fruitySpanishName,
		language.TraditionalChinese:   fruityTraditionalChineseName}
)

var (
	Fruity = nook.Villager{
		Character:   fruityCharacter,
		Personality: personality.Jock,
		Phrase:      fruityPhrase}
)
