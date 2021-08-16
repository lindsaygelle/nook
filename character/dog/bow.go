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
	bowBirthday = nook.Birthday{
		Day:   0,
		Month: time.Month(0)}
)

var (
	bowCode = nook.Code{
		Value: ""}
)

var (
	bowAmericanEnglishName = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "Bow"}

	bowCanadianFrenchName = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	bowDutchName = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	bowFrenchName = nook.Name{
		Language: language.French,
		Value:    ""}

	bowGermanName = nook.Name{
		Language: language.German,
		Value:    ""}

	bowItalianName = nook.Name{
		Language: language.Italian,
		Value:    ""}

	bowJapaneseName = nook.Name{
		Language: language.Japanese,
		Value:    "バウ"}

	bowKoreanName = nook.Name{
		Language: language.Korean,
		Value:    ""}

	bowLatinAmericanSpanishName = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	bowRussianName = nook.Name{
		Language: language.Russian,
		Value:    ""}

	bowSimplifiedChineseName = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	bowSpanishName = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	bowTraditionalChineseName = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	bowName = nook.Languages{
		language.AmericanEnglish:      bowAmericanEnglishName,
		language.CanadianFrench:       bowCanadianFrenchName,
		language.Dutch:                bowDutchName,
		language.French:               bowFrenchName,
		language.German:               bowGermanName,
		language.Italian:              bowItalianName,
		language.Japanese:             bowJapaneseName,
		language.Korean:               bowKoreanName,
		language.LatinAmericanSpanish: bowLatinAmericanSpanishName,
		language.Russian:              bowRussianName,
		language.SimplifiedChinese:    bowSimplifiedChineseName,
		language.Spanish:              bowSpanishName,
		language.TraditionalChinese:   bowTraditionalChineseName}
)

var (
	bowCharacter = nook.Character{
		Animal:   animal.Dog,
		Birthday: bowBirthday,
		Code:     bowCode,
		Key:      character.Bow,
		Gender:   gender.Male,
		Name:     bowName}
)

var (
	bowAmericanEnglishPhrase = nook.Name{
		Language: language.AmericanEnglish,
		Value:    "バウ"}

	bowCanadianFrenchPhrase = nook.Name{
		Language: language.CanadianFrench,
		Value:    ""}

	bowDutchPhrase = nook.Name{
		Language: language.Dutch,
		Value:    ""}

	bowFrenchPhrase = nook.Name{
		Language: language.French,
		Value:    ""}

	bowGermanPhrase = nook.Name{
		Language: language.German,
		Value:    ""}

	bowItalianPhrase = nook.Name{
		Language: language.Italian,
		Value:    ""}

	bowJapanesePhrase = nook.Name{
		Language: language.Japanese,
		Value:    "バウ"}

	bowKoreanPhrase = nook.Name{
		Language: language.Korean,
		Value:    ""}

	bowLatinAmericanSpanishPhrase = nook.Name{
		Language: language.LatinAmericanSpanish,
		Value:    ""}

	bowRussianPhrase = nook.Name{
		Language: language.Russian,
		Value:    ""}

	bowSimplifiedChinesePhrase = nook.Name{
		Language: language.SimplifiedChinese,
		Value:    ""}

	bowSpanishPhrase = nook.Name{
		Language: language.Spanish,
		Value:    ""}

	bowTraditionalChinesePhrase = nook.Name{
		Language: language.TraditionalChinese,
		Value:    ""}
)

var (
	bowPhrase = nook.Languages{
		language.AmericanEnglish:      bowAmericanEnglishName,
		language.CanadianFrench:       bowCanadianFrenchName,
		language.Dutch:                bowDutchName,
		language.French:               bowFrenchName,
		language.German:               bowGermanName,
		language.Italian:              bowItalianName,
		language.Japanese:             bowJapaneseName,
		language.Korean:               bowKoreanName,
		language.LatinAmericanSpanish: bowLatinAmericanSpanishName,
		language.Russian:              bowRussianName,
		language.SimplifiedChinese:    bowSimplifiedChineseName,
		language.Spanish:              bowSpanishName,
		language.TraditionalChinese:   bowTraditionalChineseName}
)

var (
	Bow = nook.Villager{
		Character:   bowCharacter,
		Personality: personality.Lazy,
		Phrase:      bowPhrase}
)
