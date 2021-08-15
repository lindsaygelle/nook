package chicken

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	kenBirthday = nook.Birthday{
		Day:   23,
		Month: time.December}
)

var (
	kenCode = nook.Code{
		Value: "chn13"}
)

var (
	kenAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Ken"}

	kenCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Kenchicken"}

	kenDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Keneiteraard"}

	kenFrenchName = nook.Name{
		Language: language.French,
		Value:    "Kenpôt pôt"}

	kenGermanName = nook.Name{
		Language: language.German,
		Value:    "Hannesgockgock"}

	kenItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Aligalliquiriqui"}

	kenJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "クロベエコッケイ"}

	kenKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "오골깜깜"}

	kenLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Pichónkirrí"}

	kenRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кенэто точно"}

	kenSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "乌骨鸡骨鸡"}

	kenSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Pichónkirrí"}

	kenTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "烏骨雞骨雞"}
)

var (
	kenName = nook.Languages{
		language.AmericanEnglish:      kenAmericanEnglishName,
		language.CanadianFrench:       kenCanadianFrenchName,
		language.Dutch:                kenDutchName,
		language.French:               kenFrenchName,
		language.German:               kenGermanName,
		language.Italian:              kenItalianName,
		language.Japanese:             kenJapaneseName,
		language.Korean:               kenKoreanName,
		language.LatinAmericanSpanish: kenLatinAmericanSpanishName,
		language.Russian:              kenRussianName,
		language.SimplifiedChinese:    kenSimplifiedChineseName,
		language.Spanish:              kenSpanishName,
		language.TraditionalChinese:   kenTraditionalChineseName}
)

var (
	kenCharacter = nook.Character{
		Animal:   Chicken,
		Birthday: kenBirthday,
		Code:     kenCode,
		Gender:   nook.Male,
		Name:     kenName}
)

var (
	kenAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "no doubt"}

	kenCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "chicken"}

	kenDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "eiteraard"}

	kenFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pôt pôt"}

	kenGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "gockgock"}

	kenItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "quiriqui"}

	kenJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "コッケイ"}

	kenKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "깜깜"}

	kenLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "kirrí"}

	kenRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "это точно"}

	kenSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "骨鸡"}

	kenSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "kirrí"}

	kenTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "骨雞"}
)

var (
	kenPhrase = nook.Languages{
		language.AmericanEnglish:      kenAmericanEnglishName,
		language.CanadianFrench:       kenCanadianFrenchName,
		language.Dutch:                kenDutchName,
		language.French:               kenFrenchName,
		language.German:               kenGermanName,
		language.Italian:              kenItalianName,
		language.Japanese:             kenJapaneseName,
		language.Korean:               kenKoreanName,
		language.LatinAmericanSpanish: kenLatinAmericanSpanishName,
		language.Russian:              kenRussianName,
		language.SimplifiedChinese:    kenSimplifiedChineseName,
		language.Spanish:              kenSpanishName,
		language.TraditionalChinese:   kenTraditionalChineseName}
)

var (
	Ken = nook.Villager{
		Character:   kenCharacter,
		Personality: nook.Smug,
		Phrase:      kenPhrase}
)
