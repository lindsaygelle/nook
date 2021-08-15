package dog

import (
	"time"

	"github.com/lindsaygelle/nook"
	"golang.org/x/text/language"
)

var (
	shepBirthday = nook.Birthday{
		Day:   24,
		Month: time.November}
)

var (
	shepCode = nook.Code{
		Value: "dog18"}
)

var (
	shepAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Shep"}

	shepCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    "Mehdipedigree"}

	shepDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Shepschapie"}

	shepFrenchName = nook.Name{
		Language: language.French,
		Value:    "Mehdipedigree"}

	shepGermanName = nook.Name{
		Language: language.German,
		Value:    "Thomasheffheff"}

	shepItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Franginoguau guau"}

	shepJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ボブワンダー"}

	shepKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "밥안보여"}

	shepLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Flecogrrruau"}

	shepRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Шепбр-р-разер"}

	shepSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "包伯汪想想"}

	shepSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Flecotufos"}

	shepTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "包伯汪想想"}
)

var (
	shepName = nook.Languages{
		language.AmericanEnglish:      shepAmericanEnglishName,
		language.CanadianFrench:       shepCanadianFrenchName,
		language.Dutch:                shepDutchName,
		language.French:               shepFrenchName,
		language.German:               shepGermanName,
		language.Italian:              shepItalianName,
		language.Japanese:             shepJapaneseName,
		language.Korean:               shepKoreanName,
		language.LatinAmericanSpanish: shepLatinAmericanSpanishName,
		language.Russian:              shepRussianName,
		language.SimplifiedChinese:    shepSimplifiedChineseName,
		language.Spanish:              shepSpanishName,
		language.TraditionalChinese:   shepTraditionalChineseName}
)

var (
	shepCharacter = nook.Character{
		Animal:   Dog,
		Birthday: shepBirthday,
		Code:     shepCode,
		Gender:   nook.Male,
		Name:     shepName}
)

var (
	shepAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "baaa manbaa baa baa"}

	shepCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    "pedigree"}

	shepDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    "schapie"}

	shepFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    "pedigree"}

	shepGermanPhrase = nook.Name{
		Language: language.German,
		Value:    "heffheff"}

	shepItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    "guau guau"}

	shepJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "ワンダー"}

	shepKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    "안보여"}

	shepLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "grrruau"}

	shepRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    "бр-р-разер"}

	shepSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "汪想想"}

	shepSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    "tufos"}

	shepTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "汪想想"}
)

var (
	shepPhrase = nook.Languages{
		language.AmericanEnglish:      shepAmericanEnglishName,
		language.CanadianFrench:       shepCanadianFrenchName,
		language.Dutch:                shepDutchName,
		language.French:               shepFrenchName,
		language.German:               shepGermanName,
		language.Italian:              shepItalianName,
		language.Japanese:             shepJapaneseName,
		language.Korean:               shepKoreanName,
		language.LatinAmericanSpanish: shepLatinAmericanSpanishName,
		language.Russian:              shepRussianName,
		language.SimplifiedChinese:    shepSimplifiedChineseName,
		language.Spanish:              shepSpanishName,
		language.TraditionalChinese:   shepTraditionalChineseName}
)

var (
	Shep = nook.Villager{
		Character:   shepCharacter,
		Personality: nook.Smug,
		Phrase:      shepPhrase}
)
