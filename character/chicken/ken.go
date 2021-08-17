package chicken

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
		Value:    "Ken"}

	kenDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Ken"}

	kenFrenchName = nook.Name{
		Language: language.French,
		Value:    "Ken"}

	kenGermanName = nook.Name{
		Language: language.German,
		Value:    "Hannes"}

	kenItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Aligalli"}

	kenJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "クロベエ"}

	kenKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "오골"}

	kenLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Pichón"}

	kenRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Кен"}

	kenSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "乌骨鸡"}

	kenSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Pichón"}

	kenTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "烏骨雞"}
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
		Animal:   animal.Chicken,
		Birthday: kenBirthday,
		Code:     kenCode,
		Key:      character.Ken,
		Gender:   gender.Male,
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
		language.AmericanEnglish:      kenAmericanEnglishPhrase,
		language.CanadianFrench:       kenCanadianFrenchPhrase,
		language.Dutch:                kenDutchPhrase,
		language.French:               kenFrenchPhrase,
		language.German:               kenGermanPhrase,
		language.Italian:              kenItalianPhrase,
		language.Japanese:             kenJapanesePhrase,
		language.Korean:               kenKoreanPhrase,
		language.LatinAmericanSpanish: kenLatinAmericanSpanishPhrase,
		language.Russian:              kenRussianPhrase,
		language.SimplifiedChinese:    kenSimplifiedChinesePhrase,
		language.Spanish:              kenSpanishPhrase,
		language.TraditionalChinese:   kenTraditionalChinesePhrase}
)

var (
	Ken = nook.Villager{
		Character:   kenCharacter,
		Personality: personality.Smug,
		Phrase:      kenPhrase}
)
