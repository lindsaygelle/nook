package dog

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
		Value:    "Mehdi"}

	shepDutchName = nook.Name{
		Language: language.Dutch,
		Value:    "Shep"}

	shepFrenchName = nook.Name{
		Language: language.French,
		Value:    "Mehdi"}

	shepGermanName = nook.Name{
		Language: language.German,
		Value:    "Thomas"}

	shepItalianName = nook.Name{
		Language: language.Italian,
		Value:    "Frangino"}

	shepJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "ボブ"}

	shepKoreanName = nook.Name{
		Language: language.Korean,
		Value:    "밥"}

	shepLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    "Fleco"}

	shepRussianName = nook.Name{
		Language: language.Russian,
		Value:    "Шеп"}

	shepSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    "包伯"}

	shepSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    "Fleco"}

	shepTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    "包伯"}
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
		Animal:   animal.Dog,
		Birthday: shepBirthday,
		Code:     shepCode,
		Key:      character.Shep,
		Gender:   gender.Male,
		Name:     shepName,
		Special:  false}
)

var (
	shepAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "baaa man"}

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
		language.AmericanEnglish:      shepAmericanEnglishPhrase,
		language.CanadianFrench:       shepCanadianFrenchPhrase,
		language.Dutch:                shepDutchPhrase,
		language.French:               shepFrenchPhrase,
		language.German:               shepGermanPhrase,
		language.Italian:              shepItalianPhrase,
		language.Japanese:             shepJapanesePhrase,
		language.Korean:               shepKoreanPhrase,
		language.LatinAmericanSpanish: shepLatinAmericanSpanishPhrase,
		language.Russian:              shepRussianPhrase,
		language.SimplifiedChinese:    shepSimplifiedChinesePhrase,
		language.Spanish:              shepSpanishPhrase,
		language.TraditionalChinese:   shepTraditionalChinesePhrase}
)

var (
	Shep = nook.Villager{
		Character:   shepCharacter,
		Personality: personality.Smug,
		Phrase:      shepPhrase}
)
