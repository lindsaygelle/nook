package duck

import (
	"time"

	"github.com/lindsaygelle/nook"
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
		Value:    "フルーティーグァバ"}

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
		Animal:   Duck,
		Birthday: fruityBirthday,
		Code:     fruityCode,
		Gender:   nook.Male,
		Name:     fruityName}
)

var (
	fruityAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    ""}

	fruityCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "N/A"}

	fruityDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "N/A"}

	fruityFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "N/A"}

	fruityGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "N/A"}

	fruityItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "N/A"}

	fruityJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "グァバ"}

	fruityKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "N/A"}

	fruityLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "N/A"}

	fruityRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "N/A"}

	fruitySimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "N/A"}

	fruitySpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "N/A"}

	fruityTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "N/A"}
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
		Personality: nook.Jock,
		Phrase:      fruityPhrase}
)
